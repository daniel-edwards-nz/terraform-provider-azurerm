package scheduledactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileDestination struct {
	FileFormats *[]FileFormat `json:"fileFormats,omitempty"`
}
