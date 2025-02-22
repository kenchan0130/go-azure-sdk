package privateclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionKeyVaultProperties struct {
	KeyName     *string                `json:"keyName,omitempty"`
	KeyState    *EncryptionKeyStatus   `json:"keyState,omitempty"`
	KeyVaultUrl *string                `json:"keyVaultUrl,omitempty"`
	KeyVersion  *string                `json:"keyVersion,omitempty"`
	VersionType *EncryptionVersionType `json:"versionType,omitempty"`
}
