/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
	"fmt"
)

// checks if the InstanceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceType{}

// InstanceType struct for InstanceType
type InstanceType struct {
	Name TypeName `json:"name"`
	Resources Resources `json:"resources"`
	// The supported storage types for the instance type
	StorageTypes []RootDiskStorageType `json:"storageTypes"`
	Prices Prices `json:"prices"`
	AdditionalProperties map[string]interface{}
}

type _InstanceType InstanceType

// NewInstanceType instantiates a new InstanceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceType(name TypeName, resources Resources, storageTypes []RootDiskStorageType, prices Prices) *InstanceType {
	this := InstanceType{}
	this.Name = name
	this.Resources = resources
	this.StorageTypes = storageTypes
	this.Prices = prices
	return &this
}

// NewInstanceTypeWithDefaults instantiates a new InstanceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceTypeWithDefaults() *InstanceType {
	this := InstanceType{}
	return &this
}

// GetName returns the Name field value
func (o *InstanceType) GetName() TypeName {
	if o == nil {
		var ret TypeName
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InstanceType) GetNameOk() (*TypeName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InstanceType) SetName(v TypeName) {
	o.Name = v
}

// GetResources returns the Resources field value
func (o *InstanceType) GetResources() Resources {
	if o == nil {
		var ret Resources
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *InstanceType) GetResourcesOk() (*Resources, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *InstanceType) SetResources(v Resources) {
	o.Resources = v
}

// GetStorageTypes returns the StorageTypes field value
// If the value is explicit nil, the zero value for []RootDiskStorageType will be returned
func (o *InstanceType) GetStorageTypes() []RootDiskStorageType {
	if o == nil {
		var ret []RootDiskStorageType
		return ret
	}

	return o.StorageTypes
}

// GetStorageTypesOk returns a tuple with the StorageTypes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceType) GetStorageTypesOk() ([]RootDiskStorageType, bool) {
	if o == nil || IsNil(o.StorageTypes) {
		return nil, false
	}
	return o.StorageTypes, true
}

// SetStorageTypes sets field value
func (o *InstanceType) SetStorageTypes(v []RootDiskStorageType) {
	o.StorageTypes = v
}

// GetPrices returns the Prices field value
func (o *InstanceType) GetPrices() Prices {
	if o == nil {
		var ret Prices
		return ret
	}

	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value
// and a boolean to check if the value has been set.
func (o *InstanceType) GetPricesOk() (*Prices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prices, true
}

// SetPrices sets field value
func (o *InstanceType) SetPrices(v Prices) {
	o.Prices = v
}

func (o InstanceType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["resources"] = o.Resources
	if o.StorageTypes != nil {
		toSerialize["storageTypes"] = o.StorageTypes
	}
	toSerialize["prices"] = o.Prices

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InstanceType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"resources",
		"storageTypes",
		"prices",
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

	varInstanceType := _InstanceType{}

	err = json.Unmarshal(data, &varInstanceType)

	if err != nil {
		return err
	}

	*o = InstanceType(varInstanceType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "resources")
		delete(additionalProperties, "storageTypes")
		delete(additionalProperties, "prices")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstanceType struct {
	value *InstanceType
	isSet bool
}

func (v NullableInstanceType) Get() *InstanceType {
	return v.value
}

func (v *NullableInstanceType) Set(val *InstanceType) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceType) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceType(val *InstanceType) *NullableInstanceType {
	return &NullableInstanceType{value: val, isSet: true}
}

func (v NullableInstanceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


