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
