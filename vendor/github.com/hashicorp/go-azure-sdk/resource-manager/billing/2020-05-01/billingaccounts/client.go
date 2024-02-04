package billingaccounts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingAccountsClient struct {
	Client *resourcemanager.Client
}

func NewBillingAccountsClientWithBaseURI(sdkApi sdkEnv.Api) (*BillingAccountsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "billingaccounts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BillingAccountsClient: %+v", err)
	}

	return &BillingAccountsClient{
		Client: client,
	}, nil
}
