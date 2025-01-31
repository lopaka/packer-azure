// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See the LICENSE file in the project root for license information.

package arm

// The intent of these types to facilitate interchange with Azure in the
// appropriate JSON format. A sample format is below.  Each parameter listed
// below corresponds to a parameter defined in the template.
//
// {
//   "storageAccountName": {
//     "value" : "my_storage_account_name"
//   },
//   "adminUserName" : {
//     "value": "admin"
//   }
// }

type TemplateParameter struct {
	Value string `json:"value"`
}

type TemplateParameters struct {
	AdminUsername      *TemplateParameter `json:"adminUsername,omitempty"`
	DnsNameForPublicIP *TemplateParameter `json:"dnsNameForPublicIP,omitempty"`
	ImageOffer         *TemplateParameter `json:"imageOffer,omitempty"`
	ImagePublisher     *TemplateParameter `json:"imagePublisher,omitempty"`
	ImageSku           *TemplateParameter `json:"imageSku,omitempty"`
	Location           *TemplateParameter `json:"location,omitempty"`
	OSDiskName         *TemplateParameter `json:"osDiskName,omitempty"`
	SshAuthorizedKey   *TemplateParameter `json:"sshAuthorizedKey,omitempty"`
	StorageAccountName *TemplateParameter `json:"storageAccountName,omitempty"`
	VMSize             *TemplateParameter `json:"vmSize,omitempty"`
	VMName             *TemplateParameter `json:"vmName,omitempty"`
}
