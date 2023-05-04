package tokens

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TokensClient struct {
	Client *resourcemanager.Client
}

func NewTokensClientWithBaseURI(api environments.Api) (*TokensClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "tokens", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TokensClient: %+v", err)
	}

	return &TokensClient{
		Client: client,
	}, nil
}
