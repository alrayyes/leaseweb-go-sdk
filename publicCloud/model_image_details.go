/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the ImageDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageDetails{}

// ImageDetails struct for ImageDetails
type ImageDetails struct {
	// imageId can be either an Operating System or a UUID in case of a Custom Image
	Id string `json:"id"`
	Name string `json:"name"`
	Family string `json:"family"`
	Flavour string `json:"flavour"`
	// Standard or Custom image
	Custom bool `json:"custom"`
	StorageSize NullableStorageSize `json:"storageSize"`
	State NullableString `json:"state"`
	// The reason in case of failure
	StateReason NullableString `json:"stateReason"`
	Region NullableRegionName `json:"region"`
	// Date when the image was created
	CreatedAt NullableTime `json:"createdAt"`
	// Date when the image was updated
	UpdatedAt NullableTime `json:"updatedAt"`
	Version string `json:"version"`
	Architecture string `json:"architecture"`
	MarketApps []string `json:"marketApps"`
	// The supported storage types for the instance type
	StorageTypes []string `json:"storageTypes"`
	AdditionalProperties map[string]interface{}
}

type _ImageDetails ImageDetails

// NewImageDetails instantiates a new ImageDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageDetails(id string, name string, family string, flavour string, custom bool, storageSize NullableStorageSize, state NullableString, stateReason NullableString, region NullableRegionName, createdAt NullableTime, updatedAt NullableTime, version string, architecture string, marketApps []string, storageTypes []string) *ImageDetails {
	this := ImageDetails{}
	this.Id = id
	this.Name = name
	this.Family = family
	this.Flavour = flavour
	this.Custom = custom
	this.StorageSize = storageSize
	this.State = state
	this.StateReason = stateReason
	this.Region = region
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Version = version
	this.Architecture = architecture
	this.MarketApps = marketApps
	this.StorageTypes = storageTypes
	return &this
}

// NewImageDetailsWithDefaults instantiates a new ImageDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageDetailsWithDefaults() *ImageDetails {
	this := ImageDetails{}
	return &this
}

// GetId returns the Id field value
func (o *ImageDetails) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ImageDetails) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ImageDetails) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ImageDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImageDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ImageDetails) SetName(v string) {
	o.Name = v
}

// GetFamily returns the Family field value
func (o *ImageDetails) GetFamily() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *ImageDetails) GetFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *ImageDetails) SetFamily(v string) {
	o.Family = v
}

// GetFlavour returns the Flavour field value
func (o *ImageDetails) GetFlavour() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Flavour
}

// GetFlavourOk returns a tuple with the Flavour field value
// and a boolean to check if the value has been set.
func (o *ImageDetails) GetFlavourOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flavour, true
}

// SetFlavour sets field value
func (o *ImageDetails) SetFlavour(v string) {
	o.Flavour = v
}

// GetCustom returns the Custom field value
func (o *ImageDetails) GetCustom() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value
// and a boolean to check if the value has been set.
func (o *ImageDetails) GetCustomOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Custom, true
}

// SetCustom sets field value
func (o *ImageDetails) SetCustom(v bool) {
	o.Custom = v
}

// GetStorageSize returns the StorageSize field value
// If the value is explicit nil, the zero value for StorageSize will be returned
func (o *ImageDetails) GetStorageSize() StorageSize {
	if o == nil || o.StorageSize.Get() == nil {
		var ret StorageSize
		return ret
	}

	return *o.StorageSize.Get()
}

// GetStorageSizeOk returns a tuple with the StorageSize field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageDetails) GetStorageSizeOk() (*StorageSize, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageSize.Get(), o.StorageSize.IsSet()
}

// SetStorageSize sets field value
func (o *ImageDetails) SetStorageSize(v StorageSize) {
	o.StorageSize.Set(&v)
}

// GetState returns the State field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ImageDetails) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}

	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageDetails) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// SetState sets field value
func (o *ImageDetails) SetState(v string) {
	o.State.Set(&v)
}

// GetStateReason returns the StateReason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ImageDetails) GetStateReason() string {
	if o == nil || o.StateReason.Get() == nil {
		var ret string
		return ret
	}

	return *o.StateReason.Get()
}

// GetStateReasonOk returns a tuple with the StateReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageDetails) GetStateReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StateReason.Get(), o.StateReason.IsSet()
}

// SetStateReason sets field value
func (o *ImageDetails) SetStateReason(v string) {
	o.StateReason.Set(&v)
}

// GetRegion returns the Region field value
// If the value is explicit nil, the zero value for RegionName will be returned
func (o *ImageDetails) GetRegion() RegionName {
	if o == nil || o.Region.Get() == nil {
		var ret RegionName
		return ret
	}

	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageDetails) GetRegionOk() (*RegionName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// SetRegion sets field value
func (o *ImageDetails) SetRegion(v RegionName) {
	o.Region.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ImageDetails) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageDetails) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *ImageDetails) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ImageDetails) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageDetails) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *ImageDetails) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// GetVersion returns the Version field value
func (o *ImageDetails) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ImageDetails) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ImageDetails) SetVersion(v string) {
	o.Version = v
}

// GetArchitecture returns the Architecture field value
func (o *ImageDetails) GetArchitecture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value
// and a boolean to check if the value has been set.
func (o *ImageDetails) GetArchitectureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Architecture, true
}

// SetArchitecture sets field value
func (o *ImageDetails) SetArchitecture(v string) {
	o.Architecture = v
}

// GetMarketApps returns the MarketApps field value
func (o *ImageDetails) GetMarketApps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MarketApps
}

// GetMarketAppsOk returns a tuple with the MarketApps field value
// and a boolean to check if the value has been set.
func (o *ImageDetails) GetMarketAppsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarketApps, true
}

// SetMarketApps sets field value
func (o *ImageDetails) SetMarketApps(v []string) {
	o.MarketApps = v
}

// GetStorageTypes returns the StorageTypes field value
func (o *ImageDetails) GetStorageTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.StorageTypes
}

// GetStorageTypesOk returns a tuple with the StorageTypes field value
// and a boolean to check if the value has been set.
func (o *ImageDetails) GetStorageTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageTypes, true
}

// SetStorageTypes sets field value
func (o *ImageDetails) SetStorageTypes(v []string) {
	o.StorageTypes = v
}

func (o ImageDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["family"] = o.Family
	toSerialize["flavour"] = o.Flavour
	toSerialize["custom"] = o.Custom
	toSerialize["storageSize"] = o.StorageSize.Get()
	toSerialize["state"] = o.State.Get()
	toSerialize["stateReason"] = o.StateReason.Get()
	toSerialize["region"] = o.Region.Get()
	toSerialize["createdAt"] = o.CreatedAt.Get()
	toSerialize["updatedAt"] = o.UpdatedAt.Get()
	toSerialize["version"] = o.Version
	toSerialize["architecture"] = o.Architecture
	toSerialize["marketApps"] = o.MarketApps
	toSerialize["storageTypes"] = o.StorageTypes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImageDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"family",
		"flavour",
		"custom",
		"storageSize",
		"state",
		"stateReason",
		"region",
		"createdAt",
		"updatedAt",
		"version",
		"architecture",
		"marketApps",
		"storageTypes",
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

	varImageDetails := _ImageDetails{}

	err = json.Unmarshal(data, &varImageDetails)

	if err != nil {
		return err
	}

	*o = ImageDetails(varImageDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "family")
		delete(additionalProperties, "flavour")
		delete(additionalProperties, "custom")
		delete(additionalProperties, "storageSize")
		delete(additionalProperties, "state")
		delete(additionalProperties, "stateReason")
		delete(additionalProperties, "region")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "version")
		delete(additionalProperties, "architecture")
		delete(additionalProperties, "marketApps")
		delete(additionalProperties, "storageTypes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImageDetails struct {
	value *ImageDetails
	isSet bool
}

func (v NullableImageDetails) Get() *ImageDetails {
	return v.value
}

func (v *NullableImageDetails) Set(val *ImageDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableImageDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableImageDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageDetails(val *ImageDetails) *NullableImageDetails {
	return &NullableImageDetails{value: val, isSet: true}
}

func (v NullableImageDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


