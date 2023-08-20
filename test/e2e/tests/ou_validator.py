from acktest import tags


class OrganizationalUnitValidator:
    def __init__(self, organizations_client):
        self.organizations_client = organizations_client

    def assert_tags(self, ou_id, expected_tags):
        response = self.organizations_client.list_tags_for_resource(ResourceId=ou_id)
        actual_tags = response['Tags']

        tags.assert_equal_without_ack_tags(
            expected=expected_tags,
            actual=actual_tags,
        )

    def assert_value(self, ou_name, expected_value):
        response = self.organizations_client.get_secret_value(SecretId=ou_name)
        actual_value = response['SecretString']

        assert actual_value == expected_value