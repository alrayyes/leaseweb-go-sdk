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

// checks if the GetReinstallImageListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReinstallImageListResult{}

// GetReinstallImageListResult struct for GetReinstallImageListResult
type GetReinstallImageListResult struct {
	Images []ImageDetails `json:"images,omitempty"`
	Metadata *Metadata `json:"_metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetReinstallImageListResult GetReinstallImageListResult

// NewGetReinstallImageListResult instantiates a new GetReinstallImageListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReinstallImageListResult() *GetReinstallImageListResult {
	this := GetReinstallImageListResult{}
	return &this
}

// NewGetReinstallImageListResultWithDefaults instantiates a new GetReinstallImageListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReinstallImageListResultWithDefaults() *GetReinstallImageListResult {
	this := GetReinstallImageListResult{}
	return &this
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *GetReinstallImageListResult) GetImages() []ImageDetails {
	if o == nil || IsNil(o.Images) {
		var ret []ImageDetails
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReinstallImageListResult) GetImagesOk() ([]ImageDetails, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *GetReinstallImageListResult) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []ImageDetails and assigns it to the Images field.
func (o *GetReinstallImageListResult) SetImages(v []ImageDetails) {
	o.Images = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetReinstallImageListResult) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReinstallImageListResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetReinstallImageListResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *GetReinstallImageListResult) SetMetadata(v Metadata) {
	o.Metadata = &v
}

func (o GetReinstallImageListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReinstallImageListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetReinstallImageListResult) UnmarshalJSON(data []byte) (err error) {
	varGetReinstallImageListResult := _GetReinstallImageListResult{}

	err = json.Unmarshal(data, &varGetReinstallImageListResult)

	if err != nil {
		return err
	}

	*o = GetReinstallImageListResult(varGetReinstallImageListResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "images")
		delete(additionalProperties, "_metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetReinstallImageListResult struct {
	value *GetReinstallImageListResult
	isSet bool
}

func (v NullableGetReinstallImageListResult) Get() *GetReinstallImageListResult {
	return v.value
}

func (v *NullableGetReinstallImageListResult) Set(val *GetReinstallImageListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReinstallImageListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReinstallImageListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReinstallImageListResult(val *GetReinstallImageListResult) *NullableGetReinstallImageListResult {
	return &NullableGetReinstallImageListResult{value: val, isSet: true}
}

func (v NullableGetReinstallImageListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReinstallImageListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

