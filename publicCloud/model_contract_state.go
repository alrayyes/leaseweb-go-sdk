/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
	"fmt"
)

// ContractState the model 'ContractState'
type ContractState string

// List of contractState
const (
	CONTRACTSTATE_ACTIVE ContractState = "ACTIVE"
	CONTRACTSTATE_DELETE_SCHEDULED ContractState = "DELETE_SCHEDULED"
)

// All allowed values of ContractState enum
var AllowedContractStateEnumValues = []ContractState{
	"ACTIVE",
	"DELETE_SCHEDULED",
}

func (v *ContractState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContractState(value)
	for _, existing := range AllowedContractStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContractState", value)
}

// NewContractStateFromValue returns a pointer to a valid ContractState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContractStateFromValue(v string) (*ContractState, error) {
	ev := ContractState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContractState: valid values are %v", v, AllowedContractStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContractState) IsValid() bool {
	for _, existing := range AllowedContractStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to contractState value
func (v ContractState) Ptr() *ContractState {
	return &v
}

type NullableContractState struct {
	value *ContractState
	isSet bool
}

func (v NullableContractState) Get() *ContractState {
	return v.value
}

func (v *NullableContractState) Set(val *ContractState) {
	v.value = val
	v.isSet = true
}

func (v NullableContractState) IsSet() bool {
	return v.isSet
}

func (v *NullableContractState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractState(val *ContractState) *NullableContractState {
	return &NullableContractState{value: val, isSet: true}
}

func (v NullableContractState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

