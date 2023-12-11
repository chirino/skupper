/*
Nexodus API

This is the Nexodus API Server.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package control_api

import (
	"encoding/json"
	"fmt"
)

// ModelsKeyUsage the model 'ModelsKeyUsage'
type ModelsKeyUsage string

// List of models.KeyUsage
const (
	UsageSigning           ModelsKeyUsage = "signing"
	UsageDigitalSignature  ModelsKeyUsage = "digital signature"
	UsageContentCommitment ModelsKeyUsage = "content commitment"
	UsageKeyEncipherment   ModelsKeyUsage = "key encipherment"
	UsageKeyAgreement      ModelsKeyUsage = "key agreement"
	UsageDataEncipherment  ModelsKeyUsage = "data encipherment"
	UsageCertSign          ModelsKeyUsage = "cert sign"
	UsageCRLSign           ModelsKeyUsage = "crl sign"
	UsageEncipherOnly      ModelsKeyUsage = "encipher only"
	UsageDecipherOnly      ModelsKeyUsage = "decipher only"
	UsageAny               ModelsKeyUsage = "any"
	UsageServerAuth        ModelsKeyUsage = "server auth"
	UsageClientAuth        ModelsKeyUsage = "client auth"
	UsageCodeSigning       ModelsKeyUsage = "code signing"
	UsageEmailProtection   ModelsKeyUsage = "email protection"
	UsageSMIME             ModelsKeyUsage = "s/mime"
	UsageIPsecEndSystem    ModelsKeyUsage = "ipsec end system"
	UsageIPsecTunnel       ModelsKeyUsage = "ipsec tunnel"
	UsageIPsecUser         ModelsKeyUsage = "ipsec user"
	UsageTimestamping      ModelsKeyUsage = "timestamping"
	UsageOCSPSigning       ModelsKeyUsage = "ocsp signing"
	UsageMicrosoftSGC      ModelsKeyUsage = "microsoft sgc"
	UsageNetscapeSGC       ModelsKeyUsage = "netscape sgc"
)

// All allowed values of ModelsKeyUsage enum
var AllowedModelsKeyUsageEnumValues = []ModelsKeyUsage{
	"signing",
	"digital signature",
	"content commitment",
	"key encipherment",
	"key agreement",
	"data encipherment",
	"cert sign",
	"crl sign",
	"encipher only",
	"decipher only",
	"any",
	"server auth",
	"client auth",
	"code signing",
	"email protection",
	"s/mime",
	"ipsec end system",
	"ipsec tunnel",
	"ipsec user",
	"timestamping",
	"ocsp signing",
	"microsoft sgc",
	"netscape sgc",
}

func (v *ModelsKeyUsage) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelsKeyUsage(value)
	for _, existing := range AllowedModelsKeyUsageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelsKeyUsage", value)
}

// NewModelsKeyUsageFromValue returns a pointer to a valid ModelsKeyUsage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelsKeyUsageFromValue(v string) (*ModelsKeyUsage, error) {
	ev := ModelsKeyUsage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelsKeyUsage: valid values are %v", v, AllowedModelsKeyUsageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelsKeyUsage) IsValid() bool {
	for _, existing := range AllowedModelsKeyUsageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to models.KeyUsage value
func (v ModelsKeyUsage) Ptr() *ModelsKeyUsage {
	return &v
}

type NullableModelsKeyUsage struct {
	value *ModelsKeyUsage
	isSet bool
}

func (v NullableModelsKeyUsage) Get() *ModelsKeyUsage {
	return v.value
}

func (v *NullableModelsKeyUsage) Set(val *ModelsKeyUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsKeyUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsKeyUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsKeyUsage(val *ModelsKeyUsage) *NullableModelsKeyUsage {
	return &NullableModelsKeyUsage{value: val, isSet: true}
}

func (v NullableModelsKeyUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsKeyUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
