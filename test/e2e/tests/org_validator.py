from acktest import tags


class OrganizationValidator:
    def __init__(self, organizations_client):
        self.organizations_client = organizations_client

    def assert_exists(self):
        response = self.organizations_client.describe_organization()

    def assert_value(self, ou_name, expected_value):
        response = self.organizations_client.get_secret_value(SecretId=ou_name)
        actual_value = response['SecretString']

        assert actual_value == expected_value