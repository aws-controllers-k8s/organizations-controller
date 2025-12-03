	// AccountID is not populated until the account creation succeeds.
	// continue describing createStatus until account finishes creating
	if r.ko.Status.AccountID == nil && r.ko.Status.CreateAccountRequestID != nil {
		r = rm.concreteResource(r.DeepCopy())
		r.ko.Status.AccountID, err = rm.getAccountID(ctx, r)
		if err != nil {
			return nil, err
		}
	}