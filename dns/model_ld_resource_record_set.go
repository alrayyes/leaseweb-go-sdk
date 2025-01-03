/*
DNS

>The base URL for this API is: **https://api.leaseweb.com/hosting/v2**

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
	"fmt"
)

// checks if the LdResourceRecordSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdResourceRecordSet{}

// LdResourceRecordSet struct for LdResourceRecordSet
type LdResourceRecordSet struct {
	// Name of the resource record set
	Name string `json:"name"`
	Type LdResourceRecordSetType `json:"type"`
	GeoContent GeoContent `json:"geoContent"`
	Ttl Ttl `json:"ttl"`
	AdditionalProperties map[string]interface{}
}

type _LdResourceRecordSet LdResourceRecordSet

// NewLdResourceRecordSet instantiates a new LdResourceRecordSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdResourceRecordSet(name string, type_ LdResourceRecordSetType, geoContent GeoContent, ttl Ttl) *LdResourceRecordSet {
	this := LdResourceRecordSet{}
	this.Name = name
	this.Type = type_
	this.GeoContent = geoContent
	this.Ttl = ttl
	return &this
}

// NewLdResourceRecordSetWithDefaults instantiates a new LdResourceRecordSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdResourceRecordSetWithDefaults() *LdResourceRecordSet {
	this := LdResourceRecordSet{}
	return &this
}

// GetName returns the Name field value
func (o *LdResourceRecordSet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LdResourceRecordSet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LdResourceRecordSet) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *LdResourceRecordSet) GetType() LdResourceRecordSetType {
	if o == nil {
		var ret LdResourceRecordSetType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LdResourceRecordSet) GetTypeOk() (*LdResourceRecordSetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LdResourceRecordSet) SetType(v LdResourceRecordSetType) {
	o.Type = v
}

// GetGeoContent returns the GeoContent field value
func (o *LdResourceRecordSet) GetGeoContent() GeoContent {
	if o == nil {
		var ret GeoContent
		return ret
	}

	return o.GeoContent
}

// GetGeoContentOk returns a tuple with the GeoContent field value
// and a boolean to check if the value has been set.
func (o *LdResourceRecordSet) GetGeoContentOk() (*GeoContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeoContent, true
}

// SetGeoContent sets field value
func (o *LdResourceRecordSet) SetGeoContent(v GeoContent) {
	o.GeoContent = v
}

// GetTtl returns the Ttl field value
func (o *LdResourceRecordSet) GetTtl() Ttl {
	if o == nil {
		var ret Ttl
		return ret
	}

	return o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value
// and a boolean to check if the value has been set.
func (o *LdResourceRecordSet) GetTtlOk() (*Ttl, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ttl, true
}

// SetTtl sets field value
func (o *LdResourceRecordSet) SetTtl(v Ttl) {
	o.Ttl = v
}

func (o LdResourceRecordSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdResourceRecordSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["geoContent"] = o.GeoContent
	toSerialize["ttl"] = o.Ttl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LdResourceRecordSet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"geoContent",
		"ttl",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLdResourceRecordSet := _LdResourceRecordSet{}

	err = json.Unmarshal(data, &varLdResourceRecordSet)

	if err != nil {
		return err
	}

	*o = LdResourceRecordSet(varLdResourceRecordSet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "geoContent")
		delete(additionalProperties, "ttl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLdResourceRecordSet struct {
	value *LdResourceRecordSet
	isSet bool
}

func (v NullableLdResourceRecordSet) Get() *LdResourceRecordSet {
	return v.value
}

func (v *NullableLdResourceRecordSet) Set(val *LdResourceRecordSet) {
	v.value = val
	v.isSet = true
}

func (v NullableLdResourceRecordSet) IsSet() bool {
	return v.isSet
}

func (v *NullableLdResourceRecordSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdResourceRecordSet(val *LdResourceRecordSet) *NullableLdResourceRecordSet {
	return &NullableLdResourceRecordSet{value: val, isSet: true}
}

func (v NullableLdResourceRecordSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdResourceRecordSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


