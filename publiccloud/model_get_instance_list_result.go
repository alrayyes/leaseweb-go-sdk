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

// checks if the GetInstanceListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInstanceListResult{}

// GetInstanceListResult struct for GetInstanceListResult
type GetInstanceListResult struct {
	Instances []Instance `json:"instances,omitempty"`
	Metadata *Metadata `json:"_metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetInstanceListResult GetInstanceListResult

// NewGetInstanceListResult instantiates a new GetInstanceListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInstanceListResult() *GetInstanceListResult {
	this := GetInstanceListResult{}
	return &this
}

// NewGetInstanceListResultWithDefaults instantiates a new GetInstanceListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInstanceListResultWithDefaults() *GetInstanceListResult {
	this := GetInstanceListResult{}
	return &this
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *GetInstanceListResult) GetInstances() []Instance {
	if o == nil || IsNil(o.Instances) {
		var ret []Instance
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceListResult) GetInstancesOk() ([]Instance, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *GetInstanceListResult) HasInstances() bool {
	if o != nil && !IsNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []Instance and assigns it to the Instances field.
func (o *GetInstanceListResult) SetInstances(v []Instance) {
	o.Instances = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetInstanceListResult) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceListResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetInstanceListResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *GetInstanceListResult) SetMetadata(v Metadata) {
	o.Metadata = &v
}

func (o GetInstanceListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInstanceListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetInstanceListResult) UnmarshalJSON(data []byte) (err error) {
	varGetInstanceListResult := _GetInstanceListResult{}

	err = json.Unmarshal(data, &varGetInstanceListResult)

	if err != nil {
		return err
	}

	*o = GetInstanceListResult(varGetInstanceListResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "instances")
		delete(additionalProperties, "_metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetInstanceListResult struct {
	value *GetInstanceListResult
	isSet bool
}

func (v NullableGetInstanceListResult) Get() *GetInstanceListResult {
	return v.value
}

func (v *NullableGetInstanceListResult) Set(val *GetInstanceListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInstanceListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInstanceListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInstanceListResult(val *GetInstanceListResult) *NullableGetInstanceListResult {
	return &NullableGetInstanceListResult{value: val, isSet: true}
}

func (v NullableGetInstanceListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInstanceListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

