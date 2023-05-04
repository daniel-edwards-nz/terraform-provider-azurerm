package services

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPrivateLinkResourceProperties struct {
	GroupId               *string                                     `json:"groupId,omitempty"`
	PrivateLinkResourceId *string                                     `json:"privateLinkResourceId,omitempty"`
	ProvisioningState     *SharedPrivateLinkResourceProvisioningState `json:"provisioningState,omitempty"`
	RequestMessage        *string                                     `json:"requestMessage,omitempty"`
	ResourceRegion        *string                                     `json:"resourceRegion,omitempty"`
	Status                *SharedPrivateLinkResourceStatus            `json:"status,omitempty"`
}
