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
from e2e.tests.ou_validator import OrganizationalUnitValidator

RESOURCE_KIND = "OrganizationalUnit"
RESOURCE_PLURAL = "organizationalunits"

DELETE_WAIT_AFTER_SECONDS = 5


@pytest.fixture(scope="module")
def simple_ou(organizations_client):
    resource_name = random_suffix_name("simple-ou", 24)

    replacements = REPLACEMENT_VALUES.copy()
    replacements["OU_NAME"] = resource_name

    # Load resource
    resource_data = load_organizations_resource(
        "organizationalunit",
        additional_replacements=replacements,
    )
    logging.debug(resource_data)

    ref = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
        resource_name, namespace="default",
    )

    # Create OU
    k8s.create_custom_resource(ref, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(ref)

    yield cr, ref

    # Get OU ID from resource
    ou_id = cr["status"]["id"]

    # Delete k8s resource
    _, deleted = k8s.delete_custom_resource(
        ref,
        period_length=DELETE_WAIT_AFTER_SECONDS,
    )
    assert deleted

    with pytest.raises(organizations_client.exceptions.OrganizationalUnitNotFoundException):
        organizations_client.describe_organizational_unit(OrganizationalUnitId=ou_id)


@service_marker
@pytest.mark.canary
class TestOrganizationalUnit:
    def test_create_delete(self, organizations_client, simple_ou):
        (res, ref) = simple_ou

        time.sleep(5)

        cr = k8s.get_resource(ref)
        assert cr is not None
        assert 'spec' in cr
        assert 'name' in cr["spec"]
        assert 'id' in cr["status"]

        ou_id = cr['status']['id']
        ou_validator = OrganizationalUnitValidator(organizations_client)
        expect_tags = {
            "key1": "value1",
        }
        ou_validator.assert_tags(ou_id, expect_tags)
