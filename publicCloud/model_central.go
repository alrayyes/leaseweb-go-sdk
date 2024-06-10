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

// checks if the Central type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Central{}

// Central struct for Central
type Central struct {
	HourlyPrice *string `json:"hourlyPrice,omitempty"`
	MonthlyPrice *string `json:"monthlyPrice,omitempty"`
}

// NewCentral instantiates a new Central object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCentral() *Central {
	this := Central{}
	return &this
}

// NewCentralWithDefaults instantiates a new Central object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCentralWithDefaults() *Central {
	this := Central{}
	return &this
}

// GetHourlyPrice returns the HourlyPrice field value if set, zero value otherwise.
func (o *Central) GetHourlyPrice() string {
	if o == nil || IsNil(o.HourlyPrice) {
		var ret string
		return ret
	}
	return *o.HourlyPrice
}

// GetHourlyPriceOk returns a tuple with the HourlyPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Central) GetHourlyPriceOk() (*string, bool) {
	if o == nil || IsNil(o.HourlyPrice) {
		return nil, false
	}
	return o.HourlyPrice, true
}

// HasHourlyPrice returns a boolean if a field has been set.
func (o *Central) HasHourlyPrice() bool {
	if o != nil && !IsNil(o.HourlyPrice) {
		return true
	}

	return false
}

// SetHourlyPrice gets a reference to the given string and assigns it to the HourlyPrice field.
func (o *Central) SetHourlyPrice(v string) {
	o.HourlyPrice = &v
}

// GetMonthlyPrice returns the MonthlyPrice field value if set, zero value otherwise.
func (o *Central) GetMonthlyPrice() string {
	if o == nil || IsNil(o.MonthlyPrice) {
		var ret string
		return ret
	}
	return *o.MonthlyPrice
}

// GetMonthlyPriceOk returns a tuple with the MonthlyPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Central) GetMonthlyPriceOk() (*string, bool) {
	if o == nil || IsNil(o.MonthlyPrice) {
		return nil, false
	}
	return o.MonthlyPrice, true
}

// HasMonthlyPrice returns a boolean if a field has been set.
func (o *Central) HasMonthlyPrice() bool {
	if o != nil && !IsNil(o.MonthlyPrice) {
		return true
	}

	return false
}

// SetMonthlyPrice gets a reference to the given string and assigns it to the MonthlyPrice field.
func (o *Central) SetMonthlyPrice(v string) {
	o.MonthlyPrice = &v
}

func (o Central) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Central) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HourlyPrice) {
		toSerialize["hourlyPrice"] = o.HourlyPrice
	}
	if !IsNil(o.MonthlyPrice) {
		toSerialize["monthlyPrice"] = o.MonthlyPrice
	}
	return toSerialize, nil
}

type NullableCentral struct {
	value *Central
	isSet bool
}

func (v NullableCentral) Get() *Central {
	return v.value
}

func (v *NullableCentral) Set(val *Central) {
	v.value = val
	v.isSet = true
}

func (v NullableCentral) IsSet() bool {
	return v.isSet
}

func (v *NullableCentral) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCentral(val *Central) *NullableCentral {
	return &NullableCentral{value: val, isSet: true}
}

func (v NullableCentral) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCentral) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

