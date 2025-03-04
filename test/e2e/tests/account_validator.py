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

from acktest import tags


class AccountValidator:
    def __init__(self, organizations_client):
        self.organizations_client = organizations_client

    def assert_tags(self, account_id, expected_tags):
        response = self.organizations_client.list_tags_for_resource(ResourceId=account_id)
        actual_tags = response['Tags']

        tags.assert_equal_without_ack_tags(
            expected=expected_tags,
            actual=actual_tags,
        )
