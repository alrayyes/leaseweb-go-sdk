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

// checks if the GetCpuMetricsResultMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCpuMetricsResultMetrics{}

// GetCpuMetricsResultMetrics struct for GetCpuMetricsResultMetrics
type GetCpuMetricsResultMetrics struct {
	CpuMetrics *CpuMetrics `json:"cpuMetrics,omitempty"`
}

// NewGetCpuMetricsResultMetrics instantiates a new GetCpuMetricsResultMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCpuMetricsResultMetrics() *GetCpuMetricsResultMetrics {
	this := GetCpuMetricsResultMetrics{}
	return &this
}

// NewGetCpuMetricsResultMetricsWithDefaults instantiates a new GetCpuMetricsResultMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCpuMetricsResultMetricsWithDefaults() *GetCpuMetricsResultMetrics {
	this := GetCpuMetricsResultMetrics{}
	return &this
}

// GetCpuMetrics returns the CpuMetrics field value if set, zero value otherwise.
func (o *GetCpuMetricsResultMetrics) GetCpuMetrics() CpuMetrics {
	if o == nil || IsNil(o.CpuMetrics) {
		var ret CpuMetrics
		return ret
	}
	return *o.CpuMetrics
}

// GetCpuMetricsOk returns a tuple with the CpuMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCpuMetricsResultMetrics) GetCpuMetricsOk() (*CpuMetrics, bool) {
	if o == nil || IsNil(o.CpuMetrics) {
		return nil, false
	}
	return o.CpuMetrics, true
}

// HasCpuMetrics returns a boolean if a field has been set.
func (o *GetCpuMetricsResultMetrics) HasCpuMetrics() bool {
	if o != nil && !IsNil(o.CpuMetrics) {
		return true
	}

	return false
}

// SetCpuMetrics gets a reference to the given CpuMetrics and assigns it to the CpuMetrics field.
func (o *GetCpuMetricsResultMetrics) SetCpuMetrics(v CpuMetrics) {
	o.CpuMetrics = &v
}

func (o GetCpuMetricsResultMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCpuMetricsResultMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CpuMetrics) {
		toSerialize["cpuMetrics"] = o.CpuMetrics
	}
	return toSerialize, nil
}

type NullableGetCpuMetricsResultMetrics struct {
	value *GetCpuMetricsResultMetrics
	isSet bool
}

func (v NullableGetCpuMetricsResultMetrics) Get() *GetCpuMetricsResultMetrics {
	return v.value
}

func (v *NullableGetCpuMetricsResultMetrics) Set(val *GetCpuMetricsResultMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCpuMetricsResultMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCpuMetricsResultMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCpuMetricsResultMetrics(val *GetCpuMetricsResultMetrics) *NullableGetCpuMetricsResultMetrics {
	return &NullableGetCpuMetricsResultMetrics{value: val, isSet: true}
}

func (v NullableGetCpuMetricsResultMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCpuMetricsResultMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

