    if ko.Status.AccountID != nil {
        tags, err := rm.fetchCurrentTags(ctx, ko.Status.AccountID)
        if err != nil {
            return nil, err
        }
        ko.Spec.Tags = FromACKTags(tags)
    }