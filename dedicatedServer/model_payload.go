/*
Leaseweb API for dedicated servers

This documents the rest api dedicatedserver provides.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedServer

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Payload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Payload{}

// Payload struct for Payload
type Payload struct {
	FileserverBaseUrl *string `json:"fileserverBaseUrl,omitempty"`
	JobType string `json:"jobType"`
	Pop *string `json:"pop,omitempty"`
	PowerCycle *bool `json:"powerCycle,omitempty"`
	IsUnassignedServer *bool `json:"isUnassignedServer,omitempty"`
	// Id of the server
	ServerId *string `json:"serverId,omitempty"`
}

type _Payload Payload

// NewPayload instantiates a new Payload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayload(jobType string) *Payload {
	this := Payload{}
	this.JobType = jobType
	return &this
}

// NewPayloadWithDefaults instantiates a new Payload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayloadWithDefaults() *Payload {
	this := Payload{}
	return &this
}

// GetFileserverBaseUrl returns the FileserverBaseUrl field value if set, zero value otherwise.
func (o *Payload) GetFileserverBaseUrl() string {
	if o == nil || IsNil(o.FileserverBaseUrl) {
		var ret string
		return ret
	}
	return *o.FileserverBaseUrl
}

// GetFileserverBaseUrlOk returns a tuple with the FileserverBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetFileserverBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FileserverBaseUrl) {
		return nil, false
	}
	return o.FileserverBaseUrl, true
}

// HasFileserverBaseUrl returns a boolean if a field has been set.
func (o *Payload) HasFileserverBaseUrl() bool {
	if o != nil && !IsNil(o.FileserverBaseUrl) {
		return true
	}

	return false
}

// SetFileserverBaseUrl gets a reference to the given string and assigns it to the FileserverBaseUrl field.
func (o *Payload) SetFileserverBaseUrl(v string) {
	o.FileserverBaseUrl = &v
}

// GetJobType returns the JobType field value
func (o *Payload) GetJobType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value
// and a boolean to check if the value has been set.
func (o *Payload) GetJobTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobType, true
}

// SetJobType sets field value
func (o *Payload) SetJobType(v string) {
	o.JobType = v
}

// GetPop returns the Pop field value if set, zero value otherwise.
func (o *Payload) GetPop() string {
	if o == nil || IsNil(o.Pop) {
		var ret string
		return ret
	}
	return *o.Pop
}

// GetPopOk returns a tuple with the Pop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetPopOk() (*string, bool) {
	if o == nil || IsNil(o.Pop) {
		return nil, false
	}
	return o.Pop, true
}

// HasPop returns a boolean if a field has been set.
func (o *Payload) HasPop() bool {
	if o != nil && !IsNil(o.Pop) {
		return true
	}

	return false
}

// SetPop gets a reference to the given string and assigns it to the Pop field.
func (o *Payload) SetPop(v string) {
	o.Pop = &v
}

// GetPowerCycle returns the PowerCycle field value if set, zero value otherwise.
func (o *Payload) GetPowerCycle() bool {
	if o == nil || IsNil(o.PowerCycle) {
		var ret bool
		return ret
	}
	return *o.PowerCycle
}

// GetPowerCycleOk returns a tuple with the PowerCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetPowerCycleOk() (*bool, bool) {
	if o == nil || IsNil(o.PowerCycle) {
		return nil, false
	}
	return o.PowerCycle, true
}

// HasPowerCycle returns a boolean if a field has been set.
func (o *Payload) HasPowerCycle() bool {
	if o != nil && !IsNil(o.PowerCycle) {
		return true
	}

	return false
}

// SetPowerCycle gets a reference to the given bool and assigns it to the PowerCycle field.
func (o *Payload) SetPowerCycle(v bool) {
	o.PowerCycle = &v
}

// GetIsUnassignedServer returns the IsUnassignedServer field value if set, zero value otherwise.
func (o *Payload) GetIsUnassignedServer() bool {
	if o == nil || IsNil(o.IsUnassignedServer) {
		var ret bool
		return ret
	}
	return *o.IsUnassignedServer
}

// GetIsUnassignedServerOk returns a tuple with the IsUnassignedServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetIsUnassignedServerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUnassignedServer) {
		return nil, false
	}
	return o.IsUnassignedServer, true
}

// HasIsUnassignedServer returns a boolean if a field has been set.
func (o *Payload) HasIsUnassignedServer() bool {
	if o != nil && !IsNil(o.IsUnassignedServer) {
		return true
	}

	return false
}

// SetIsUnassignedServer gets a reference to the given bool and assigns it to the IsUnassignedServer field.
func (o *Payload) SetIsUnassignedServer(v bool) {
	o.IsUnassignedServer = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *Payload) GetServerId() string {
	if o == nil || IsNil(o.ServerId) {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetServerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *Payload) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *Payload) SetServerId(v string) {
	o.ServerId = &v
}

func (o Payload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Payload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileserverBaseUrl) {
		toSerialize["fileserverBaseUrl"] = o.FileserverBaseUrl
	}
	toSerialize["jobType"] = o.JobType
	if !IsNil(o.Pop) {
		toSerialize["pop"] = o.Pop
	}
	if !IsNil(o.PowerCycle) {
		toSerialize["powerCycle"] = o.PowerCycle
	}
	if !IsNil(o.IsUnassignedServer) {
		toSerialize["isUnassignedServer"] = o.IsUnassignedServer
	}
	if !IsNil(o.ServerId) {
		toSerialize["serverId"] = o.ServerId
	}
	return toSerialize, nil
}

func (o *Payload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"jobType",
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

	varPayload := _Payload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPayload)

	if err != nil {
		return err
	}

	*o = Payload(varPayload)

	return err
}

type NullablePayload struct {
	value *Payload
	isSet bool
}

func (v NullablePayload) Get() *Payload {
	return v.value
}

func (v *NullablePayload) Set(val *Payload) {
	v.value = val
	v.isSet = true
}

func (v NullablePayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayload(val *Payload) *NullablePayload {
	return &NullablePayload{value: val, isSet: true}
}

func (v NullablePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

