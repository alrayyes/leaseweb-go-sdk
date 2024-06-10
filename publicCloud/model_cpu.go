/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
)

// checks if the Cpu type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cpu{}

// Cpu Number of cores
type Cpu struct {
	Value *int32 `json:"value,omitempty"`
	Unit *string `json:"unit,omitempty"`
}

// NewCpu instantiates a new Cpu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpu() *Cpu {
	this := Cpu{}
	return &this
}

// NewCpuWithDefaults instantiates a new Cpu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpuWithDefaults() *Cpu {
	this := Cpu{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Cpu) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cpu) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Cpu) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *Cpu) SetValue(v int32) {
	o.Value = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Cpu) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cpu) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Cpu) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *Cpu) SetUnit(v string) {
	o.Unit = &v
}

func (o Cpu) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cpu) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	return toSerialize, nil
}

type NullableCpu struct {
	value *Cpu
	isSet bool
}

func (v NullableCpu) Get() *Cpu {
	return v.value
}

func (v *NullableCpu) Set(val *Cpu) {
	v.value = val
	v.isSet = true
}

func (v NullableCpu) IsSet() bool {
	return v.isSet
}

func (v *NullableCpu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpu(val *Cpu) *NullableCpu {
	return &NullableCpu{value: val, isSet: true}
}

func (v NullableCpu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

