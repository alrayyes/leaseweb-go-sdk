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

// checks if the GetUpdateInstanceTypeListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUpdateInstanceTypeListResult{}

// GetUpdateInstanceTypeListResult struct for GetUpdateInstanceTypeListResult
type GetUpdateInstanceTypeListResult struct {
	Metadata *Metadata `json:"_metadata,omitempty"`
	InstanceTypes []UpdateInstanceType `json:"instanceTypes,omitempty"`
}

// NewGetUpdateInstanceTypeListResult instantiates a new GetUpdateInstanceTypeListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUpdateInstanceTypeListResult() *GetUpdateInstanceTypeListResult {
	this := GetUpdateInstanceTypeListResult{}
	return &this
}

// NewGetUpdateInstanceTypeListResultWithDefaults instantiates a new GetUpdateInstanceTypeListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUpdateInstanceTypeListResultWithDefaults() *GetUpdateInstanceTypeListResult {
	this := GetUpdateInstanceTypeListResult{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetUpdateInstanceTypeListResult) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUpdateInstanceTypeListResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetUpdateInstanceTypeListResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *GetUpdateInstanceTypeListResult) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetInstanceTypes returns the InstanceTypes field value if set, zero value otherwise.
func (o *GetUpdateInstanceTypeListResult) GetInstanceTypes() []UpdateInstanceType {
	if o == nil || IsNil(o.InstanceTypes) {
		var ret []UpdateInstanceType
		return ret
	}
	return o.InstanceTypes
}

// GetInstanceTypesOk returns a tuple with the InstanceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUpdateInstanceTypeListResult) GetInstanceTypesOk() ([]UpdateInstanceType, bool) {
	if o == nil || IsNil(o.InstanceTypes) {
		return nil, false
	}
	return o.InstanceTypes, true
}

// HasInstanceTypes returns a boolean if a field has been set.
func (o *GetUpdateInstanceTypeListResult) HasInstanceTypes() bool {
	if o != nil && !IsNil(o.InstanceTypes) {
		return true
	}

	return false
}

// SetInstanceTypes gets a reference to the given []UpdateInstanceType and assigns it to the InstanceTypes field.
func (o *GetUpdateInstanceTypeListResult) SetInstanceTypes(v []UpdateInstanceType) {
	o.InstanceTypes = v
}

func (o GetUpdateInstanceTypeListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUpdateInstanceTypeListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}
	if !IsNil(o.InstanceTypes) {
		toSerialize["instanceTypes"] = o.InstanceTypes
	}
	return toSerialize, nil
}

type NullableGetUpdateInstanceTypeListResult struct {
	value *GetUpdateInstanceTypeListResult
	isSet bool
}

func (v NullableGetUpdateInstanceTypeListResult) Get() *GetUpdateInstanceTypeListResult {
	return v.value
}

func (v *NullableGetUpdateInstanceTypeListResult) Set(val *GetUpdateInstanceTypeListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUpdateInstanceTypeListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUpdateInstanceTypeListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUpdateInstanceTypeListResult(val *GetUpdateInstanceTypeListResult) *NullableGetUpdateInstanceTypeListResult {
	return &NullableGetUpdateInstanceTypeListResult{value: val, isSet: true}
}

func (v NullableGetUpdateInstanceTypeListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUpdateInstanceTypeListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

