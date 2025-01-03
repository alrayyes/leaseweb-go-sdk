/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
	"fmt"
)

// VpsPackType Vps package type
type VpsPackType string

// List of vpsPackType
const (
	VPSPACKTYPE__1 VpsPackType = "Leaseweb VPS 1"
	VPSPACKTYPE__2 VpsPackType = "Leaseweb VPS 2"
	VPSPACKTYPE__3 VpsPackType = "Leaseweb VPS 3"
	VPSPACKTYPE__4 VpsPackType = "Leaseweb VPS 4"
	VPSPACKTYPE__5 VpsPackType = "Leaseweb VPS 5"
	VPSPACKTYPE__6 VpsPackType = "Leaseweb VPS 6"
)

// All allowed values of VpsPackType enum
var AllowedVpsPackTypeEnumValues = []VpsPackType{
	"Leaseweb VPS 1",
	"Leaseweb VPS 2",
	"Leaseweb VPS 3",
	"Leaseweb VPS 4",
	"Leaseweb VPS 5",
	"Leaseweb VPS 6",
}

func (v *VpsPackType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VpsPackType(value)
	for _, existing := range AllowedVpsPackTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VpsPackType", value)
}

// NewVpsPackTypeFromValue returns a pointer to a valid VpsPackType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVpsPackTypeFromValue(v string) (*VpsPackType, error) {
	ev := VpsPackType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VpsPackType: valid values are %v", v, AllowedVpsPackTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VpsPackType) IsValid() bool {
	for _, existing := range AllowedVpsPackTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vpsPackType value
func (v VpsPackType) Ptr() *VpsPackType {
	return &v
}

type NullableVpsPackType struct {
	value *VpsPackType
	isSet bool
}

func (v NullableVpsPackType) Get() *VpsPackType {
	return v.value
}

func (v *NullableVpsPackType) Set(val *VpsPackType) {
	v.value = val
	v.isSet = true
}

func (v NullableVpsPackType) IsSet() bool {
	return v.isSet
}

func (v *NullableVpsPackType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpsPackType(val *VpsPackType) *NullableVpsPackType {
	return &NullableVpsPackType{value: val, isSet: true}
}

func (v NullableVpsPackType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpsPackType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

