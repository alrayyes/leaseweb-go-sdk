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

// InstanceState The instance's current state
type InstanceState string

// List of instanceState
const (
	CREATING InstanceState = "CREATING"
	DESTROYED InstanceState = "DESTROYED"
	DESTROYING InstanceState = "DESTROYING"
	FAILED InstanceState = "FAILED"
	RUNNING InstanceState = "RUNNING"
	STARTING InstanceState = "STARTING"
	STOPPED InstanceState = "STOPPED"
	STOPPING InstanceState = "STOPPING"
	UNKNOWN InstanceState = "UNKNOWN"
)

// All allowed values of InstanceState enum
var AllowedInstanceStateEnumValues = []InstanceState{
	"CREATING",
	"DESTROYED",
	"DESTROYING",
	"FAILED",
	"RUNNING",
	"STARTING",
	"STOPPED",
	"STOPPING",
	"UNKNOWN",
}

func (v *InstanceState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InstanceState(value)
	for _, existing := range AllowedInstanceStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InstanceState", value)
}

// NewInstanceStateFromValue returns a pointer to a valid InstanceState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInstanceStateFromValue(v string) (*InstanceState, error) {
	ev := InstanceState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InstanceState: valid values are %v", v, AllowedInstanceStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InstanceState) IsValid() bool {
	for _, existing := range AllowedInstanceStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to instanceState value
func (v InstanceState) Ptr() *InstanceState {
	return &v
}

type NullableInstanceState struct {
	value *InstanceState
	isSet bool
}

func (v NullableInstanceState) Get() *InstanceState {
	return v.value
}

func (v *NullableInstanceState) Set(val *InstanceState) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceState) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceState(val *InstanceState) *NullableInstanceState {
	return &NullableInstanceState{value: val, isSet: true}
}

func (v NullableInstanceState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

