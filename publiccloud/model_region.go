/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
	"fmt"
)

// checks if the Region type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Region{}

// Region struct for Region
type Region struct {
	Name RegionName `json:"name"`
	// The city where the region is located.
	Location string `json:"location"`
	AdditionalProperties map[string]interface{}
}

type _Region Region

// NewRegion instantiates a new Region object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegion(name RegionName, location string) *Region {
	this := Region{}
	this.Name = name
	this.Location = location
	return &this
}

// NewRegionWithDefaults instantiates a new Region object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionWithDefaults() *Region {
	this := Region{}
	return &this
}

// GetName returns the Name field value
func (o *Region) GetName() RegionName {
	if o == nil {
		var ret RegionName
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Region) GetNameOk() (*RegionName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Region) SetName(v RegionName) {
	o.Name = v
}

// GetLocation returns the Location field value
func (o *Region) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Region) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Region) SetLocation(v string) {
	o.Location = v
}

func (o Region) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Region) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["location"] = o.Location

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Region) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"location",
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

	varRegion := _Region{}

	err = json.Unmarshal(data, &varRegion)

	if err != nil {
		return err
	}

	*o = Region(varRegion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "location")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegion struct {
	value *Region
	isSet bool
}

func (v NullableRegion) Get() *Region {
	return v.value
}

func (v *NullableRegion) Set(val *Region) {
	v.value = val
	v.isSet = true
}

func (v NullableRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegion(val *Region) *NullableRegion {
	return &NullableRegion{value: val, isSet: true}
}

func (v NullableRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

