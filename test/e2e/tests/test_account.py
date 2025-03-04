# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
# 	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Integration tests for the Organizations API.
"""

import logging
import pytest
import time
from acktest import tags
from acktest.k8s import resource as k8s
from acktest.resources import random_suffix_name
from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_organizations_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e.tests.account_validator import AccountValidator

RESOURCE_KIND = "Account"
RESOURCE_PLURAL = "accounts"

# Accounts take a bit to go into a SUCCEEDED state
CREATE_WAIT_SECONDS = 120

# Accounts can take a bit to become Suspended/Closed and as such be removed.
DELETE_WAIT_AFTER_SECONDS = 60

def organization_exists(organizations_client):
    try:
        _ = organizations_client.describe_organization()
        return True
    except organizations_client.exceptions.AWSOrganizationsNotInUseException:
        return False
    except Exception as e:
        assert False, f"describe_organizations failed with exception {str(e)}"



@pytest.fixture(scope="module")
def basic_account(organizations_client):
    resource_name = random_suffix_name("acktest", 12)

    org_exists = organization_exists(organizations_client)
    if False == org_exists:
        print("org doesn't exist, creating...")
        try:
            _ = organizations_client.create_organization(FeatureSet="ALL")
        except:
            assert False, f"create_organization failed with exception {str(e)}"

    list_root_resp = organizations_client.list_roots()
    root_id = list_root_resp["Roots"][0]["Id"]

    replacements = REPLACEMENT_VALUES.copy()
    replacements["ACCOUNT_NAME"] = resource_name
    replacements["ACCOUNT_EMAIL"] = f"{resource_name}@example.com"

    # Load resource
    resource_data = load_organizations_resource(
        "account",
        additional_replacements=replacements,
    )
    logging.debug(resource_data)

    ref = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
        resource_name, namespace="default",
    )

    # Create Account
    k8s.create_custom_resource(ref, resource_data)

    # Wait for the account to be in a SUCCEEDED state before we get the resource out of k8s
    time.sleep(CREATE_WAIT_SECONDS)

    cr = k8s.wait_resource_consumed_by_controller(ref)

    yield cr, ref

    # Get Account ID from resource
    account_id = cr["status"]["accountID"]

    # Delete k8s resource
    _, deleted = k8s.delete_custom_resource(
        ref,
        period_length=DELETE_WAIT_AFTER_SECONDS,
    )
    assert deleted

    with pytest.raises(organizations_client.exceptions.AccountAlreadyClosedException):
        organizations_client.close_account(AccountId=account_id)

    # Delete Organization if we created one
    if False == org_exists:
        try:
            organizations_client.delete_organization()
        except Exception as e:
            assert False, f"deleting organisation {root_id} failed with exception {str(e)}"



@service_marker
@pytest.mark.canary
class TestAccount:
    def test_create_close(self, organizations_client, basic_account):
        (res, ref) = basic_account

        time.sleep(5)

        cr = k8s.get_resource(ref)
        assert cr is not None
        assert 'spec' in cr
        assert 'name' in cr["spec"]
        assert 'accountID' in cr["status"]

        account_id = cr['status']['accountID']
        account_validator = AccountValidator(organizations_client)
        expect_tags = {
            "key1": "value1",
        }
        account_validator.assert_tags(account_id, expect_tags)
