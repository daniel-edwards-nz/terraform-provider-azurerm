package scheduledqueryrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Schedule struct {
	FrequencyInMinutes  int64 `json:"frequencyInMinutes"`
	TimeWindowInMinutes int64 `json:"timeWindowInMinutes"`
}
