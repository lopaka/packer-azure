// Copyright (c) Microsoft Open Technologies, Inc.
// All Rights Reserved.
// Licensed under the Apache License, Version 2.0.
// See License.txt in the project root for license information.

package arm

import (
	"fmt"

	"github.com/Azure/packer-azure/packer/builder/azure/common/utils"
)

const (
	TempNameAlphabet = "0123456789bcdfghjklmnpqrstvwxyz"
)

type TempName struct {
	ComputeName       string
	DeploymentName    string
	ResourceGroupName string
	OSDiskName        string
}

func NewTempName() *TempName {
	tempName := &TempName{}

	suffix := utils.RandomString(TempNameAlphabet, 10)
	tempName.ComputeName = fmt.Sprintf("pkrvm%s", suffix)
	tempName.DeploymentName = fmt.Sprintf("pkrdp%s", suffix)
	tempName.OSDiskName = fmt.Sprintf("pkros%s", suffix)
	tempName.ResourceGroupName = fmt.Sprintf("packer-Resource-Group-%s", suffix)

	return tempName
}