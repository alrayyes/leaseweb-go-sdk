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

// checks if the Collection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Collection{}

// Collection Link to the resource record sets collection
type Collection struct {
	// Hyperlink to the resource record sets collection location
	Href *string `json:"href,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Collection Collection

// NewCollection instantiates a new Collection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollection() *Collection {
	this := Collection{}
	return &this
}

// NewCollectionWithDefaults instantiates a new Collection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionWithDefaults() *Collection {
	this := Collection{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Collection) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collection) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Collection) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Collection) SetHref(v string) {
	o.Href = &v
}

func (o Collection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Collection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Collection) UnmarshalJSON(data []byte) (err error) {
	varCollection := _Collection{}

	err = json.Unmarshal(data, &varCollection)

	if err != nil {
		return err
	}

	*o = Collection(varCollection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollection struct {
	value *Collection
	isSet bool
}

func (v NullableCollection) Get() *Collection {
	return v.value
}

func (v *NullableCollection) Set(val *Collection) {
	v.value = val
	v.isSet = true
}

func (v NullableCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollection(val *Collection) *NullableCollection {
	return &NullableCollection{value: val, isSet: true}
}

func (v NullableCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


