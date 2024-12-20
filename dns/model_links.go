/*
DNS

>The base URL for this API is: **https://api.leaseweb.com/hosting/v2**

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the Links type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Links{}

// Links Links to related resource locations
type Links struct {
	Self *Self `json:"self,omitempty"`
	Parent *Parent `json:"parent,omitempty"`
	ValidateSet *ValidateSet `json:"validateSet,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Links Links

// NewLinks instantiates a new Links object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinks() *Links {
	this := Links{}
	return &this
}

// NewLinksWithDefaults instantiates a new Links object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksWithDefaults() *Links {
	this := Links{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *Links) GetSelf() Self {
	if o == nil || IsNil(o.Self) {
		var ret Self
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetSelfOk() (*Self, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *Links) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Self and assigns it to the Self field.
func (o *Links) SetSelf(v Self) {
	o.Self = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *Links) GetParent() Parent {
	if o == nil || IsNil(o.Parent) {
		var ret Parent
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetParentOk() (*Parent, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *Links) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given Parent and assigns it to the Parent field.
func (o *Links) SetParent(v Parent) {
	o.Parent = &v
}

// GetValidateSet returns the ValidateSet field value if set, zero value otherwise.
func (o *Links) GetValidateSet() ValidateSet {
	if o == nil || IsNil(o.ValidateSet) {
		var ret ValidateSet
		return ret
	}
	return *o.ValidateSet
}

// GetValidateSetOk returns a tuple with the ValidateSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetValidateSetOk() (*ValidateSet, bool) {
	if o == nil || IsNil(o.ValidateSet) {
		return nil, false
	}
	return o.ValidateSet, true
}

// HasValidateSet returns a boolean if a field has been set.
func (o *Links) HasValidateSet() bool {
	if o != nil && !IsNil(o.ValidateSet) {
		return true
	}

	return false
}

// SetValidateSet gets a reference to the given ValidateSet and assigns it to the ValidateSet field.
func (o *Links) SetValidateSet(v ValidateSet) {
	o.ValidateSet = &v
}

func (o Links) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Links) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.ValidateSet) {
		toSerialize["validateSet"] = o.ValidateSet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Links) UnmarshalJSON(data []byte) (err error) {
	varLinks := _Links{}

	err = json.Unmarshal(data, &varLinks)

	if err != nil {
		return err
	}

	*o = Links(varLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "validateSet")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinks struct {
	value *Links
	isSet bool
}

func (v NullableLinks) Get() *Links {
	return v.value
}

func (v *NullableLinks) Set(val *Links) {
	v.value = val
	v.isSet = true
}

func (v NullableLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinks(val *Links) *NullableLinks {
	return &NullableLinks{value: val, isSet: true}
}

func (v NullableLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


