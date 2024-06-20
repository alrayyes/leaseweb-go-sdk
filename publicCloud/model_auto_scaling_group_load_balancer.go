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
)

// checks if the AutoScalingGroupLoadBalancer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoScalingGroupLoadBalancer{}

// AutoScalingGroupLoadBalancer struct for AutoScalingGroupLoadBalancer
type AutoScalingGroupLoadBalancer struct {
	// The customers ID who owns the load balancer
	CustomerId *string `json:"customerId,omitempty"`
	// The load balancer unique identifier
	Id *string `json:"id,omitempty"`
	// Load balancer type
	Type *string `json:"type,omitempty"`
	Resources *InstanceResources `json:"resources,omitempty"`
	// The region where the load balancer was launched into
	Region *string `json:"region,omitempty"`
	// The identifying name set to the load balancer
	Reference *string `json:"reference,omitempty"`
	State *InstanceState `json:"state,omitempty"`
	Ips []Ip `json:"ips,omitempty"`
	Contract *Contract `json:"contract,omitempty"`
	// Date and time when the instance was started for the first time, right after launching it
	StartedAt *time.Time `json:"startedAt,omitempty"`
	Configuration NullableLoadBalancerConfiguration `json:"configuration,omitempty"`
	PrivateNetwork NullablePrivateNetwork `json:"privateNetwork,omitempty"`
}

// NewAutoScalingGroupLoadBalancer instantiates a new AutoScalingGroupLoadBalancer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoScalingGroupLoadBalancer() *AutoScalingGroupLoadBalancer {
	this := AutoScalingGroupLoadBalancer{}
	return &this
}

// NewAutoScalingGroupLoadBalancerWithDefaults instantiates a new AutoScalingGroupLoadBalancer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoScalingGroupLoadBalancerWithDefaults() *AutoScalingGroupLoadBalancer {
	this := AutoScalingGroupLoadBalancer{}
	return &this
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *AutoScalingGroupLoadBalancer) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScalingGroupLoadBalancer) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *AutoScalingGroupLoadBalancer) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *AutoScalingGroupLoadBalancer) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutoScalingGroupLoadBalancer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScalingGroupLoadBalancer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutoScalingGroupLoadBalancer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AutoScalingGroupLoadBalancer) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AutoScalingGroupLoadBalancer) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScalingGroupLoadBalancer) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AutoScalingGroupLoadBalancer) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AutoScalingGroupLoadBalancer) SetType(v string) {
	o.Type = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *AutoScalingGroupLoadBalancer) GetResources() InstanceResources {
	if o == nil || IsNil(o.Resources) {
		var ret InstanceResources
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScalingGroupLoadBalancer) GetResourcesOk() (*InstanceResources, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *AutoScalingGroupLoadBalancer) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given InstanceResources and assigns it to the Resources field.
func (o *AutoScalingGroupLoadBalancer) SetResources(v InstanceResources) {
	o.Resources = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AutoScalingGroupLoadBalancer) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScalingGroupLoadBalancer) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AutoScalingGroupLoadBalancer) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AutoScalingGroupLoadBalancer) SetRegion(v string) {
	o.Region = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *AutoScalingGroupLoadBalancer) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScalingGroupLoadBalancer) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *AutoScalingGroupLoadBalancer) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *AutoScalingGroupLoadBalancer) SetReference(v string) {
	o.Reference = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AutoScalingGroupLoadBalancer) GetState() InstanceState {
	if o == nil || IsNil(o.State) {
		var ret InstanceState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScalingGroupLoadBalancer) GetStateOk() (*InstanceState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AutoScalingGroupLoadBalancer) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given InstanceState and assigns it to the State field.
func (o *AutoScalingGroupLoadBalancer) SetState(v InstanceState) {
	o.State = &v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *AutoScalingGroupLoadBalancer) GetIps() []Ip {
	if o == nil || IsNil(o.Ips) {
		var ret []Ip
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScalingGroupLoadBalancer) GetIpsOk() ([]Ip, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *AutoScalingGroupLoadBalancer) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given []Ip and assigns it to the Ips field.
func (o *AutoScalingGroupLoadBalancer) SetIps(v []Ip) {
	o.Ips = v
}

// GetContract returns the Contract field value if set, zero value otherwise.
func (o *AutoScalingGroupLoadBalancer) GetContract() Contract {
	if o == nil || IsNil(o.Contract) {
		var ret Contract
		return ret
	}
	return *o.Contract
}

// GetContractOk returns a tuple with the Contract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScalingGroupLoadBalancer) GetContractOk() (*Contract, bool) {
	if o == nil || IsNil(o.Contract) {
		return nil, false
	}
	return o.Contract, true
}

// HasContract returns a boolean if a field has been set.
func (o *AutoScalingGroupLoadBalancer) HasContract() bool {
	if o != nil && !IsNil(o.Contract) {
		return true
	}

	return false
}

// SetContract gets a reference to the given Contract and assigns it to the Contract field.
func (o *AutoScalingGroupLoadBalancer) SetContract(v Contract) {
	o.Contract = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *AutoScalingGroupLoadBalancer) GetStartedAt() time.Time {
	if o == nil || IsNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScalingGroupLoadBalancer) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *AutoScalingGroupLoadBalancer) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *AutoScalingGroupLoadBalancer) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoScalingGroupLoadBalancer) GetConfiguration() LoadBalancerConfiguration {
	if o == nil || IsNil(o.Configuration.Get()) {
		var ret LoadBalancerConfiguration
		return ret
	}
	return *o.Configuration.Get()
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoScalingGroupLoadBalancer) GetConfigurationOk() (*LoadBalancerConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Configuration.Get(), o.Configuration.IsSet()
}

// HasConfiguration returns a boolean if a field has been set.
func (o *AutoScalingGroupLoadBalancer) HasConfiguration() bool {
	if o != nil && o.Configuration.IsSet() {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given NullableLoadBalancerConfiguration and assigns it to the Configuration field.
func (o *AutoScalingGroupLoadBalancer) SetConfiguration(v LoadBalancerConfiguration) {
	o.Configuration.Set(&v)
}
// SetConfigurationNil sets the value for Configuration to be an explicit nil
func (o *AutoScalingGroupLoadBalancer) SetConfigurationNil() {
	o.Configuration.Set(nil)
}

// UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
func (o *AutoScalingGroupLoadBalancer) UnsetConfiguration() {
	o.Configuration.Unset()
}

// GetPrivateNetwork returns the PrivateNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoScalingGroupLoadBalancer) GetPrivateNetwork() PrivateNetwork {
	if o == nil || IsNil(o.PrivateNetwork.Get()) {
		var ret PrivateNetwork
		return ret
	}
	return *o.PrivateNetwork.Get()
}

// GetPrivateNetworkOk returns a tuple with the PrivateNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoScalingGroupLoadBalancer) GetPrivateNetworkOk() (*PrivateNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateNetwork.Get(), o.PrivateNetwork.IsSet()
}

// HasPrivateNetwork returns a boolean if a field has been set.
func (o *AutoScalingGroupLoadBalancer) HasPrivateNetwork() bool {
	if o != nil && o.PrivateNetwork.IsSet() {
		return true
	}

	return false
}

// SetPrivateNetwork gets a reference to the given NullablePrivateNetwork and assigns it to the PrivateNetwork field.
func (o *AutoScalingGroupLoadBalancer) SetPrivateNetwork(v PrivateNetwork) {
	o.PrivateNetwork.Set(&v)
}
// SetPrivateNetworkNil sets the value for PrivateNetwork to be an explicit nil
func (o *AutoScalingGroupLoadBalancer) SetPrivateNetworkNil() {
	o.PrivateNetwork.Set(nil)
}

// UnsetPrivateNetwork ensures that no value is present for PrivateNetwork, not even an explicit nil
func (o *AutoScalingGroupLoadBalancer) UnsetPrivateNetwork() {
	o.PrivateNetwork.Unset()
}

func (o AutoScalingGroupLoadBalancer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoScalingGroupLoadBalancer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Ips) {
		toSerialize["ips"] = o.Ips
	}
	if !IsNil(o.Contract) {
		toSerialize["contract"] = o.Contract
	}
	if !IsNil(o.StartedAt) {
		toSerialize["startedAt"] = o.StartedAt
	}
	if o.Configuration.IsSet() {
		toSerialize["configuration"] = o.Configuration.Get()
	}
	if o.PrivateNetwork.IsSet() {
		toSerialize["privateNetwork"] = o.PrivateNetwork.Get()
	}
	return toSerialize, nil
}

type NullableAutoScalingGroupLoadBalancer struct {
	value *AutoScalingGroupLoadBalancer
	isSet bool
}

func (v NullableAutoScalingGroupLoadBalancer) Get() *AutoScalingGroupLoadBalancer {
	return v.value
}

func (v *NullableAutoScalingGroupLoadBalancer) Set(val *AutoScalingGroupLoadBalancer) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoScalingGroupLoadBalancer) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoScalingGroupLoadBalancer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoScalingGroupLoadBalancer(val *AutoScalingGroupLoadBalancer) *NullableAutoScalingGroupLoadBalancer {
	return &NullableAutoScalingGroupLoadBalancer{value: val, isSet: true}
}

func (v NullableAutoScalingGroupLoadBalancer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoScalingGroupLoadBalancer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

