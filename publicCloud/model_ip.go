/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Ip type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ip{}

// Ip struct for Ip
type Ip struct {
	// Ip Address
	Ip string `json:"ip"`
	// The number of leading bits in the IP address
	PrefixLength string `json:"prefixLength"`
	Version IpVersion `json:"version"`
	// Whether or not the IP has been nulled
	NullRouted bool `json:"nullRouted"`
	MainIp bool `json:"mainIp"`
	NetworkType NetworkType `json:"networkType"`
	ReverseLookup NullableString `json:"reverseLookup"`
}

type _Ip Ip

// NewIp instantiates a new Ip object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIp(ip string, prefixLength string, version IpVersion, nullRouted bool, mainIp bool, networkType NetworkType, reverseLookup NullableString) *Ip {
	this := Ip{}
	this.Ip = ip
	this.PrefixLength = prefixLength
	this.Version = version
	this.NullRouted = nullRouted
	this.MainIp = mainIp
	this.NetworkType = networkType
	this.ReverseLookup = reverseLookup
	return &this
}

// NewIpWithDefaults instantiates a new Ip object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpWithDefaults() *Ip {
	this := Ip{}
	return &this
}

// GetIp returns the Ip field value
func (o *Ip) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *Ip) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *Ip) SetIp(v string) {
	o.Ip = v
}

// GetPrefixLength returns the PrefixLength field value
func (o *Ip) GetPrefixLength() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrefixLength
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value
// and a boolean to check if the value has been set.
func (o *Ip) GetPrefixLengthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrefixLength, true
}

// SetPrefixLength sets field value
func (o *Ip) SetPrefixLength(v string) {
	o.PrefixLength = v
}

// GetVersion returns the Version field value
func (o *Ip) GetVersion() IpVersion {
	if o == nil {
		var ret IpVersion
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Ip) GetVersionOk() (*IpVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Ip) SetVersion(v IpVersion) {
	o.Version = v
}

// GetNullRouted returns the NullRouted field value
func (o *Ip) GetNullRouted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NullRouted
}

// GetNullRoutedOk returns a tuple with the NullRouted field value
// and a boolean to check if the value has been set.
func (o *Ip) GetNullRoutedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NullRouted, true
}

// SetNullRouted sets field value
func (o *Ip) SetNullRouted(v bool) {
	o.NullRouted = v
}

// GetMainIp returns the MainIp field value
func (o *Ip) GetMainIp() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MainIp
}

// GetMainIpOk returns a tuple with the MainIp field value
// and a boolean to check if the value has been set.
func (o *Ip) GetMainIpOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MainIp, true
}

// SetMainIp sets field value
func (o *Ip) SetMainIp(v bool) {
	o.MainIp = v
}

// GetNetworkType returns the NetworkType field value
func (o *Ip) GetNetworkType() NetworkType {
	if o == nil {
		var ret NetworkType
		return ret
	}

	return o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value
// and a boolean to check if the value has been set.
func (o *Ip) GetNetworkTypeOk() (*NetworkType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkType, true
}

// SetNetworkType sets field value
func (o *Ip) SetNetworkType(v NetworkType) {
	o.NetworkType = v
}

// GetReverseLookup returns the ReverseLookup field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Ip) GetReverseLookup() string {
	if o == nil || o.ReverseLookup.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReverseLookup.Get()
}

// GetReverseLookupOk returns a tuple with the ReverseLookup field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ip) GetReverseLookupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReverseLookup.Get(), o.ReverseLookup.IsSet()
}

// SetReverseLookup sets field value
func (o *Ip) SetReverseLookup(v string) {
	o.ReverseLookup.Set(&v)
}

func (o Ip) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ip) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ip"] = o.Ip
	toSerialize["prefixLength"] = o.PrefixLength
	toSerialize["version"] = o.Version
	toSerialize["nullRouted"] = o.NullRouted
	toSerialize["mainIp"] = o.MainIp
	toSerialize["networkType"] = o.NetworkType
	toSerialize["reverseLookup"] = o.ReverseLookup.Get()
	return toSerialize, nil
}

func (o *Ip) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ip",
		"prefixLength",
		"version",
		"nullRouted",
		"mainIp",
		"networkType",
		"reverseLookup",
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

	varIp := _Ip{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIp)

	if err != nil {
		return err
	}

	*o = Ip(varIp)

	return err
}

type NullableIp struct {
	value *Ip
	isSet bool
}

func (v NullableIp) Get() *Ip {
	return v.value
}

func (v *NullableIp) Set(val *Ip) {
	v.value = val
	v.isSet = true
}

func (v NullableIp) IsSet() bool {
	return v.isSet
}

func (v *NullableIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIp(val *Ip) *NullableIp {
	return &NullableIp{value: val, isSet: true}
}

func (v NullableIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


