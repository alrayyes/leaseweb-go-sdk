/*
Leaseweb API for dedicated servers

This documents the rest api dedicatedserver provides.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedserver

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateNetworkEquipmentCredentialOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkEquipmentCredentialOpts{}

// UpdateNetworkEquipmentCredentialOpts struct for UpdateNetworkEquipmentCredentialOpts
type UpdateNetworkEquipmentCredentialOpts struct {
	// The password for the credentials
	Password string `json:"password"`
	AdditionalProperties map[string]interface{}
}

type _UpdateNetworkEquipmentCredentialOpts UpdateNetworkEquipmentCredentialOpts

// NewUpdateNetworkEquipmentCredentialOpts instantiates a new UpdateNetworkEquipmentCredentialOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkEquipmentCredentialOpts(password string) *UpdateNetworkEquipmentCredentialOpts {
	this := UpdateNetworkEquipmentCredentialOpts{}
	this.Password = password
	return &this
}

// NewUpdateNetworkEquipmentCredentialOptsWithDefaults instantiates a new UpdateNetworkEquipmentCredentialOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkEquipmentCredentialOptsWithDefaults() *UpdateNetworkEquipmentCredentialOpts {
	this := UpdateNetworkEquipmentCredentialOpts{}
	return &this
}

// GetPassword returns the Password field value
func (o *UpdateNetworkEquipmentCredentialOpts) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkEquipmentCredentialOpts) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UpdateNetworkEquipmentCredentialOpts) SetPassword(v string) {
	o.Password = v
}

func (o UpdateNetworkEquipmentCredentialOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkEquipmentCredentialOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateNetworkEquipmentCredentialOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateNetworkEquipmentCredentialOpts := _UpdateNetworkEquipmentCredentialOpts{}

	err = json.Unmarshal(data, &varUpdateNetworkEquipmentCredentialOpts)

	if err != nil {
		return err
	}

	*o = UpdateNetworkEquipmentCredentialOpts(varUpdateNetworkEquipmentCredentialOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateNetworkEquipmentCredentialOpts struct {
	value *UpdateNetworkEquipmentCredentialOpts
	isSet bool
}

func (v NullableUpdateNetworkEquipmentCredentialOpts) Get() *UpdateNetworkEquipmentCredentialOpts {
	return v.value
}

func (v *NullableUpdateNetworkEquipmentCredentialOpts) Set(val *UpdateNetworkEquipmentCredentialOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkEquipmentCredentialOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkEquipmentCredentialOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkEquipmentCredentialOpts(val *UpdateNetworkEquipmentCredentialOpts) *NullableUpdateNetworkEquipmentCredentialOpts {
	return &NullableUpdateNetworkEquipmentCredentialOpts{value: val, isSet: true}
}

func (v NullableUpdateNetworkEquipmentCredentialOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkEquipmentCredentialOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


