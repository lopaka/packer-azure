// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See the LICENSE file in the project root for license information.

package arm

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestTemplateParametersShouldHaveExpectedKeys(t *testing.T) {
	params := TemplateParameters{
		AdminUsername:      &TemplateParameter{"sentinel"},
		DnsNameForPublicIP: &TemplateParameter{"sentinel"},
		ImageOffer:         &TemplateParameter{"sentinel"},
		ImagePublisher:     &TemplateParameter{"sentinel"},
		ImageSku:           &TemplateParameter{"sentinel"},
		Location:           &TemplateParameter{"sentinel"},
		OSDiskName:         &TemplateParameter{"sentinel"},
		SshAuthorizedKey:   &TemplateParameter{"sentinel"},
		StorageAccountName: &TemplateParameter{"sentinel"},
		VMName:             &TemplateParameter{"sentinel"},
		VMSize:             &TemplateParameter{"sentinel"},
	}

	bs, err := json.Marshal(params)
	if err != nil {
		t.Fail()
	}

	var doc map[string]*json.RawMessage
	err = json.Unmarshal(bs, &doc)

	if err != nil {
		t.Fail()
	}

	expectedKeys := []string{
		"adminUsername",
		"dnsNameForPublicIP",
		"imageOffer",
		"imagePublisher",
		"imageSku",
		"location",
		"osDiskName",
		"sshAuthorizedKey",
		"storageAccountName",
		"vmSize",
		"vmName",
	}

	for _, expectedKey := range expectedKeys {
		_, containsKey := doc[expectedKey]
		if containsKey == false {
			t.Fatalf("Expected template parameters to contain the key value '%s', but it did not!", expectedKey)
		}
	}
}

func TestParameterValuesShouldBeSet(t *testing.T) {
	params := TemplateParameters{
		AdminUsername:      &TemplateParameter{"adminusername00"},
		DnsNameForPublicIP: &TemplateParameter{"dnsnameforpublicip00"},
		ImageOffer:         &TemplateParameter{"imageoffer00"},
		ImagePublisher:     &TemplateParameter{"imagepublisher00"},
		ImageSku:           &TemplateParameter{"imagesku00"},
		Location:           &TemplateParameter{"location00"},
		OSDiskName:         &TemplateParameter{"osdiskname00"},
		SshAuthorizedKey:   &TemplateParameter{"sshauthorizedkey00"},
		StorageAccountName: &TemplateParameter{"storageaccountname00"},
		VMName:             &TemplateParameter{"vmname00"},
		VMSize:             &TemplateParameter{"vmsize00"},
	}

	bs, err := json.Marshal(params)
	if err != nil {
		t.Fail()
	}

	var doc map[string]map[string]interface{}
	err = json.Unmarshal(bs, &doc)

	if err != nil {
		t.Fail()
	}

	for k, v := range doc {
		var expectedValue = fmt.Sprintf("%s00", strings.ToLower(k))
		var actualValue, exists = v["value"]
		if exists != true {
			t.Errorf("Expected to find a 'value' key under '%s', but it was missing!", k)
		}

		if expectedValue != actualValue {
			t.Errorf("Expected '%s', but actual was '%s'!", expectedValue, actualValue)
		}
	}
}

func TestEmptyValuesShouldBeOmitted(t *testing.T) {
	params := TemplateParameters{
		AdminUsername: &TemplateParameter{"adminusername00"},
	}

	bs, err := json.Marshal(params)
	if err != nil {
		t.Fail()
	}

	var doc map[string]map[string]interface{}
	err = json.Unmarshal(bs, &doc)

	if err != nil {
		t.Fail()
	}

	if len(doc) != 1 {
		t.Errorf("Failed to omit empty template parameters from the JSON document!")
		t.Errorf("doc=%+v", doc)
		t.Fail()
	}
}
