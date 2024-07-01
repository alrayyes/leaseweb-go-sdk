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
	"bytes"
	"fmt"
)

// checks if the InstanceBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceBase{}

// InstanceBase struct for InstanceBase
type InstanceBase struct {
	// The unique identifier of the instance
	Id string `json:"id"`
	Type InstanceTypeName `json:"type"`
	Resources Resources `json:"resources"`
	// The region in which the instance was launched
	Region string `json:"region"`
	// The identifying name set to the instance
	Reference NullableString `json:"reference"`
	// Date and time when the instance was started for the first time, right after launching it
	StartedAt NullableTime `json:"startedAt"`
	// Market App ID that must be installed into the instance
	MarketAppId NullableString `json:"marketAppId"`
	State State `json:"state"`
	// The product type
	ProductType string `json:"productType"`
	HasPublicIpV4 bool `json:"hasPublicIpV4"`
	IncludesPrivateNetwork bool `json:"hasPrivateNetwork"`
	// The root disk's size in GB. Must be at least 5 GB for Linux and FreeBSD instances and 50 GB for Windows instances
	RootDiskSize int32 `json:"rootDiskSize"`
	RootDiskStorageType RootDiskStorageType `json:"rootDiskStorageType"`
	Contract Contract `json:"contract"`
	AutoScalingGroup NullableAutoScalingGroupDetails `json:"autoScalingGroup"`
}

type _InstanceBase InstanceBase

// NewInstanceBase instantiates a new InstanceBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceBase(id string, type_ InstanceTypeName, resources Resources, region string, reference NullableString, startedAt NullableTime, marketAppId NullableString, state State, productType string, hasPublicIpV4 bool, includesPrivateNetwork bool, rootDiskSize int32, rootDiskStorageType RootDiskStorageType, contract Contract, autoScalingGroup NullableAutoScalingGroupDetails) *InstanceBase {
	this := InstanceBase{}
	this.Id = id
	this.Type = type_
	this.Resources = resources
	this.Region = region
	this.Reference = reference
	this.StartedAt = startedAt
	this.MarketAppId = marketAppId
	this.State = state
	this.ProductType = productType
	this.HasPublicIpV4 = hasPublicIpV4
	this.IncludesPrivateNetwork = includesPrivateNetwork
	this.RootDiskSize = rootDiskSize
	this.RootDiskStorageType = rootDiskStorageType
	this.Contract = contract
	this.AutoScalingGroup = autoScalingGroup
	return &this
}

// NewInstanceBaseWithDefaults instantiates a new InstanceBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceBaseWithDefaults() *InstanceBase {
	this := InstanceBase{}
	return &this
}

// GetId returns the Id field value
func (o *InstanceBase) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InstanceBase) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InstanceBase) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *InstanceBase) GetType() InstanceTypeName {
	if o == nil {
		var ret InstanceTypeName
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InstanceBase) GetTypeOk() (*InstanceTypeName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InstanceBase) SetType(v InstanceTypeName) {
	o.Type = v
}

// GetResources returns the Resources field value
func (o *InstanceBase) GetResources() Resources {
	if o == nil {
		var ret Resources
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *InstanceBase) GetResourcesOk() (*Resources, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *InstanceBase) SetResources(v Resources) {
	o.Resources = v
}

// GetRegion returns the Region field value
func (o *InstanceBase) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *InstanceBase) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *InstanceBase) SetRegion(v string) {
	o.Region = v
}

// GetReference returns the Reference field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InstanceBase) GetReference() string {
	if o == nil || o.Reference.Get() == nil {
		var ret string
		return ret
	}

	return *o.Reference.Get()
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceBase) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reference.Get(), o.Reference.IsSet()
}

// SetReference sets field value
func (o *InstanceBase) SetReference(v string) {
	o.Reference.Set(&v)
}

// GetStartedAt returns the StartedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *InstanceBase) GetStartedAt() time.Time {
	if o == nil || o.StartedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceBase) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// SetStartedAt sets field value
func (o *InstanceBase) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}

// GetMarketAppId returns the MarketAppId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InstanceBase) GetMarketAppId() string {
	if o == nil || o.MarketAppId.Get() == nil {
		var ret string
		return ret
	}

	return *o.MarketAppId.Get()
}

// GetMarketAppIdOk returns a tuple with the MarketAppId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceBase) GetMarketAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarketAppId.Get(), o.MarketAppId.IsSet()
}

// SetMarketAppId sets field value
func (o *InstanceBase) SetMarketAppId(v string) {
	o.MarketAppId.Set(&v)
}

// GetState returns the State field value
func (o *InstanceBase) GetState() State {
	if o == nil {
		var ret State
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *InstanceBase) GetStateOk() (*State, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *InstanceBase) SetState(v State) {
	o.State = v
}

// GetProductType returns the ProductType field value
func (o *InstanceBase) GetProductType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *InstanceBase) GetProductTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *InstanceBase) SetProductType(v string) {
	o.ProductType = v
}

// GetHasPublicIpV4 returns the HasPublicIpV4 field value
func (o *InstanceBase) GetHasPublicIpV4() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPublicIpV4
}

// GetHasPublicIpV4Ok returns a tuple with the HasPublicIpV4 field value
// and a boolean to check if the value has been set.
func (o *InstanceBase) GetHasPublicIpV4Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPublicIpV4, true
}

// SetHasPublicIpV4 sets field value
func (o *InstanceBase) SetHasPublicIpV4(v bool) {
	o.HasPublicIpV4 = v
}

// GetIncludesPrivateNetwork returns the IncludesPrivateNetwork field value
func (o *InstanceBase) GetIncludesPrivateNetwork() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludesPrivateNetwork
}

// GetIncludesPrivateNetworkOk returns a tuple with the IncludesPrivateNetwork field value
// and a boolean to check if the value has been set.
func (o *InstanceBase) GetIncludesPrivateNetworkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludesPrivateNetwork, true
}

// SetIncludesPrivateNetwork sets field value
func (o *InstanceBase) SetIncludesPrivateNetwork(v bool) {
	o.IncludesPrivateNetwork = v
}

// GetRootDiskSize returns the RootDiskSize field value
func (o *InstanceBase) GetRootDiskSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RootDiskSize
}

// GetRootDiskSizeOk returns a tuple with the RootDiskSize field value
// and a boolean to check if the value has been set.
func (o *InstanceBase) GetRootDiskSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootDiskSize, true
}

// SetRootDiskSize sets field value
func (o *InstanceBase) SetRootDiskSize(v int32) {
	o.RootDiskSize = v
}

// GetRootDiskStorageType returns the RootDiskStorageType field value
func (o *InstanceBase) GetRootDiskStorageType() RootDiskStorageType {
	if o == nil {
		var ret RootDiskStorageType
		return ret
	}

	return o.RootDiskStorageType
}

// GetRootDiskStorageTypeOk returns a tuple with the RootDiskStorageType field value
// and a boolean to check if the value has been set.
func (o *InstanceBase) GetRootDiskStorageTypeOk() (*RootDiskStorageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootDiskStorageType, true
}

// SetRootDiskStorageType sets field value
func (o *InstanceBase) SetRootDiskStorageType(v RootDiskStorageType) {
	o.RootDiskStorageType = v
}

// GetContract returns the Contract field value
func (o *InstanceBase) GetContract() Contract {
	if o == nil {
		var ret Contract
		return ret
	}

	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value
// and a boolean to check if the value has been set.
func (o *InstanceBase) GetContractOk() (*Contract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contract, true
}

// SetContract sets field value
func (o *InstanceBase) SetContract(v Contract) {
	o.Contract = v
}

// GetAutoScalingGroup returns the AutoScalingGroup field value
// If the value is explicit nil, the zero value for AutoScalingGroupDetails will be returned
func (o *InstanceBase) GetAutoScalingGroup() AutoScalingGroupDetails {
	if o == nil || o.AutoScalingGroup.Get() == nil {
		var ret AutoScalingGroupDetails
		return ret
	}

	return *o.AutoScalingGroup.Get()
}

// GetAutoScalingGroupOk returns a tuple with the AutoScalingGroup field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceBase) GetAutoScalingGroupOk() (*AutoScalingGroupDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoScalingGroup.Get(), o.AutoScalingGroup.IsSet()
}

// SetAutoScalingGroup sets field value
func (o *InstanceBase) SetAutoScalingGroup(v AutoScalingGroupDetails) {
	o.AutoScalingGroup.Set(&v)
}

func (o InstanceBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["resources"] = o.Resources
	toSerialize["region"] = o.Region
	toSerialize["reference"] = o.Reference.Get()
	toSerialize["startedAt"] = o.StartedAt.Get()
	toSerialize["marketAppId"] = o.MarketAppId.Get()
	toSerialize["state"] = o.State
	toSerialize["productType"] = o.ProductType
	toSerialize["hasPublicIpV4"] = o.HasPublicIpV4
	toSerialize["hasPrivateNetwork"] = o.IncludesPrivateNetwork
	toSerialize["rootDiskSize"] = o.RootDiskSize
	toSerialize["rootDiskStorageType"] = o.RootDiskStorageType
	toSerialize["contract"] = o.Contract
	toSerialize["autoScalingGroup"] = o.AutoScalingGroup.Get()
	return toSerialize, nil
}

func (o *InstanceBase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"resources",
		"region",
		"reference",
		"startedAt",
		"marketAppId",
		"state",
		"productType",
		"hasPublicIpV4",
		"hasPrivateNetwork",
		"rootDiskSize",
		"rootDiskStorageType",
		"contract",
		"autoScalingGroup",
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

	varInstanceBase := _InstanceBase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInstanceBase)

	if err != nil {
		return err
	}

	*o = InstanceBase(varInstanceBase)

	return err
}

type NullableInstanceBase struct {
	value *InstanceBase
	isSet bool
}

func (v NullableInstanceBase) Get() *InstanceBase {
	return v.value
}

func (v *NullableInstanceBase) Set(val *InstanceBase) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceBase) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceBase(val *InstanceBase) *NullableInstanceBase {
	return &NullableInstanceBase{value: val, isSet: true}
}

func (v NullableInstanceBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


