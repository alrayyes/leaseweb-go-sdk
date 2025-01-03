/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
)

// checks if the Summary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Summary{}

// Summary struct for Summary
type Summary struct {
	DOWN_PUBLIC *TrafficMetricSummary `json:"DOWN_PUBLIC,omitempty"`
	UP_PUBLIC *TrafficMetricSummary `json:"UP_PUBLIC,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Summary Summary

// NewSummary instantiates a new Summary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSummary() *Summary {
	this := Summary{}
	return &this
}

// NewSummaryWithDefaults instantiates a new Summary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSummaryWithDefaults() *Summary {
	this := Summary{}
	return &this
}

// GetDOWN_PUBLIC returns the DOWN_PUBLIC field value if set, zero value otherwise.
func (o *Summary) GetDOWN_PUBLIC() TrafficMetricSummary {
	if o == nil || IsNil(o.DOWN_PUBLIC) {
		var ret TrafficMetricSummary
		return ret
	}
	return *o.DOWN_PUBLIC
}

// GetDOWN_PUBLICOk returns a tuple with the DOWN_PUBLIC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetDOWN_PUBLICOk() (*TrafficMetricSummary, bool) {
	if o == nil || IsNil(o.DOWN_PUBLIC) {
		return nil, false
	}
	return o.DOWN_PUBLIC, true
}

// HasDOWN_PUBLIC returns a boolean if a field has been set.
func (o *Summary) HasDOWN_PUBLIC() bool {
	if o != nil && !IsNil(o.DOWN_PUBLIC) {
		return true
	}

	return false
}

// SetDOWN_PUBLIC gets a reference to the given TrafficMetricSummary and assigns it to the DOWN_PUBLIC field.
func (o *Summary) SetDOWN_PUBLIC(v TrafficMetricSummary) {
	o.DOWN_PUBLIC = &v
}

// GetUP_PUBLIC returns the UP_PUBLIC field value if set, zero value otherwise.
func (o *Summary) GetUP_PUBLIC() TrafficMetricSummary {
	if o == nil || IsNil(o.UP_PUBLIC) {
		var ret TrafficMetricSummary
		return ret
	}
	return *o.UP_PUBLIC
}

// GetUP_PUBLICOk returns a tuple with the UP_PUBLIC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetUP_PUBLICOk() (*TrafficMetricSummary, bool) {
	if o == nil || IsNil(o.UP_PUBLIC) {
		return nil, false
	}
	return o.UP_PUBLIC, true
}

// HasUP_PUBLIC returns a boolean if a field has been set.
func (o *Summary) HasUP_PUBLIC() bool {
	if o != nil && !IsNil(o.UP_PUBLIC) {
		return true
	}

	return false
}

// SetUP_PUBLIC gets a reference to the given TrafficMetricSummary and assigns it to the UP_PUBLIC field.
func (o *Summary) SetUP_PUBLIC(v TrafficMetricSummary) {
	o.UP_PUBLIC = &v
}

func (o Summary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Summary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DOWN_PUBLIC) {
		toSerialize["DOWN_PUBLIC"] = o.DOWN_PUBLIC
	}
	if !IsNil(o.UP_PUBLIC) {
		toSerialize["UP_PUBLIC"] = o.UP_PUBLIC
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Summary) UnmarshalJSON(data []byte) (err error) {
	varSummary := _Summary{}

	err = json.Unmarshal(data, &varSummary)

	if err != nil {
		return err
	}

	*o = Summary(varSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "DOWN_PUBLIC")
		delete(additionalProperties, "UP_PUBLIC")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSummary struct {
	value *Summary
	isSet bool
}

func (v NullableSummary) Get() *Summary {
	return v.value
}

func (v *NullableSummary) Set(val *Summary) {
	v.value = val
	v.isSet = true
}

func (v NullableSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummary(val *Summary) *NullableSummary {
	return &NullableSummary{value: val, isSet: true}
}

func (v NullableSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


