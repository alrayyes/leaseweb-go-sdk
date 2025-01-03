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

// checks if the TrafficMetricSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrafficMetricSummary{}

// TrafficMetricSummary struct for TrafficMetricSummary
type TrafficMetricSummary struct {
	// Average incoming traffic based on the amount of grouped data points, in human-readable format (KB, MB, GB)
	Average *string `json:"average,omitempty"`
	// Expected incoming traffic given the average times the amount of days between the `from` and `to` dates, in human-readable format (KB, MB, GB)
	Expected *string `json:"expected,omitempty"`
	// Total incoming traffic, in human-readable format (KB, MB, GB)
	Total *string `json:"total,omitempty"`
	Peak *Peak `json:"peak,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TrafficMetricSummary TrafficMetricSummary

// NewTrafficMetricSummary instantiates a new TrafficMetricSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrafficMetricSummary() *TrafficMetricSummary {
	this := TrafficMetricSummary{}
	return &this
}

// NewTrafficMetricSummaryWithDefaults instantiates a new TrafficMetricSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrafficMetricSummaryWithDefaults() *TrafficMetricSummary {
	this := TrafficMetricSummary{}
	return &this
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *TrafficMetricSummary) GetAverage() string {
	if o == nil || IsNil(o.Average) {
		var ret string
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficMetricSummary) GetAverageOk() (*string, bool) {
	if o == nil || IsNil(o.Average) {
		return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *TrafficMetricSummary) HasAverage() bool {
	if o != nil && !IsNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given string and assigns it to the Average field.
func (o *TrafficMetricSummary) SetAverage(v string) {
	o.Average = &v
}

// GetExpected returns the Expected field value if set, zero value otherwise.
func (o *TrafficMetricSummary) GetExpected() string {
	if o == nil || IsNil(o.Expected) {
		var ret string
		return ret
	}
	return *o.Expected
}

// GetExpectedOk returns a tuple with the Expected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficMetricSummary) GetExpectedOk() (*string, bool) {
	if o == nil || IsNil(o.Expected) {
		return nil, false
	}
	return o.Expected, true
}

// HasExpected returns a boolean if a field has been set.
func (o *TrafficMetricSummary) HasExpected() bool {
	if o != nil && !IsNil(o.Expected) {
		return true
	}

	return false
}

// SetExpected gets a reference to the given string and assigns it to the Expected field.
func (o *TrafficMetricSummary) SetExpected(v string) {
	o.Expected = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *TrafficMetricSummary) GetTotal() string {
	if o == nil || IsNil(o.Total) {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficMetricSummary) GetTotalOk() (*string, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *TrafficMetricSummary) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *TrafficMetricSummary) SetTotal(v string) {
	o.Total = &v
}

// GetPeak returns the Peak field value if set, zero value otherwise.
func (o *TrafficMetricSummary) GetPeak() Peak {
	if o == nil || IsNil(o.Peak) {
		var ret Peak
		return ret
	}
	return *o.Peak
}

// GetPeakOk returns a tuple with the Peak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficMetricSummary) GetPeakOk() (*Peak, bool) {
	if o == nil || IsNil(o.Peak) {
		return nil, false
	}
	return o.Peak, true
}

// HasPeak returns a boolean if a field has been set.
func (o *TrafficMetricSummary) HasPeak() bool {
	if o != nil && !IsNil(o.Peak) {
		return true
	}

	return false
}

// SetPeak gets a reference to the given Peak and assigns it to the Peak field.
func (o *TrafficMetricSummary) SetPeak(v Peak) {
	o.Peak = &v
}

func (o TrafficMetricSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrafficMetricSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	if !IsNil(o.Expected) {
		toSerialize["expected"] = o.Expected
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Peak) {
		toSerialize["peak"] = o.Peak
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TrafficMetricSummary) UnmarshalJSON(data []byte) (err error) {
	varTrafficMetricSummary := _TrafficMetricSummary{}

	err = json.Unmarshal(data, &varTrafficMetricSummary)

	if err != nil {
		return err
	}

	*o = TrafficMetricSummary(varTrafficMetricSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "average")
		delete(additionalProperties, "expected")
		delete(additionalProperties, "total")
		delete(additionalProperties, "peak")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTrafficMetricSummary struct {
	value *TrafficMetricSummary
	isSet bool
}

func (v NullableTrafficMetricSummary) Get() *TrafficMetricSummary {
	return v.value
}

func (v *NullableTrafficMetricSummary) Set(val *TrafficMetricSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficMetricSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficMetricSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficMetricSummary(val *TrafficMetricSummary) *NullableTrafficMetricSummary {
	return &NullableTrafficMetricSummary{value: val, isSet: true}
}

func (v NullableTrafficMetricSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficMetricSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


