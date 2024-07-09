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

// checks if the IpDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpDetails{}

// IpDetails struct for IpDetails
type IpDetails struct {
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
	Ddos NullableDdos `json:"ddos"`
}

type _IpDetails IpDetails

// NewIpDetails instantiates a new IpDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpDetails(ip string, prefixLength string, version IpVersion, nullRouted bool, mainIp bool, networkType NetworkType, reverseLookup NullableString, ddos NullableDdos) *IpDetails {
	this := IpDetails{}
	this.Ip = ip
	this.PrefixLength = prefixLength
	this.Version = version
	this.NullRouted = nullRouted
	this.MainIp = mainIp
	this.NetworkType = networkType
	this.ReverseLookup = reverseLookup
	this.Ddos = ddos
	return &this
}

// NewIpDetailsWithDefaults instantiates a new IpDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpDetailsWithDefaults() *IpDetails {
	this := IpDetails{}
	return &this
}

// GetIp returns the Ip field value
func (o *IpDetails) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *IpDetails) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *IpDetails) SetIp(v string) {
	o.Ip = v
}

// GetPrefixLength returns the PrefixLength field value
func (o *IpDetails) GetPrefixLength() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrefixLength
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value
// and a boolean to check if the value has been set.
func (o *IpDetails) GetPrefixLengthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrefixLength, true
}

// SetPrefixLength sets field value
func (o *IpDetails) SetPrefixLength(v string) {
	o.PrefixLength = v
}

// GetVersion returns the Version field value
func (o *IpDetails) GetVersion() IpVersion {
	if o == nil {
		var ret IpVersion
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *IpDetails) GetVersionOk() (*IpVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *IpDetails) SetVersion(v IpVersion) {
	o.Version = v
}

// GetNullRouted returns the NullRouted field value
func (o *IpDetails) GetNullRouted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NullRouted
}

// GetNullRoutedOk returns a tuple with the NullRouted field value
// and a boolean to check if the value has been set.
func (o *IpDetails) GetNullRoutedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NullRouted, true
}

// SetNullRouted sets field value
func (o *IpDetails) SetNullRouted(v bool) {
	o.NullRouted = v
}

// GetMainIp returns the MainIp field value
func (o *IpDetails) GetMainIp() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MainIp
}

// GetMainIpOk returns a tuple with the MainIp field value
// and a boolean to check if the value has been set.
func (o *IpDetails) GetMainIpOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MainIp, true
}

// SetMainIp sets field value
func (o *IpDetails) SetMainIp(v bool) {
	o.MainIp = v
}

// GetNetworkType returns the NetworkType field value
func (o *IpDetails) GetNetworkType() NetworkType {
	if o == nil {
		var ret NetworkType
		return ret
	}

	return o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value
// and a boolean to check if the value has been set.
func (o *IpDetails) GetNetworkTypeOk() (*NetworkType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkType, true
}

// SetNetworkType sets field value
func (o *IpDetails) SetNetworkType(v NetworkType) {
	o.NetworkType = v
}

// GetReverseLookup returns the ReverseLookup field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IpDetails) GetReverseLookup() string {
	if o == nil || o.ReverseLookup.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReverseLookup.Get()
}

// GetReverseLookupOk returns a tuple with the ReverseLookup field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpDetails) GetReverseLookupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReverseLookup.Get(), o.ReverseLookup.IsSet()
}

// SetReverseLookup sets field value
func (o *IpDetails) SetReverseLookup(v string) {
	o.ReverseLookup.Set(&v)
}

// GetDdos returns the Ddos field value
// If the value is explicit nil, the zero value for Ddos will be returned
func (o *IpDetails) GetDdos() Ddos {
	if o == nil || o.Ddos.Get() == nil {
		var ret Ddos
		return ret
	}

	return *o.Ddos.Get()
}

// GetDdosOk returns a tuple with the Ddos field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpDetails) GetDdosOk() (*Ddos, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ddos.Get(), o.Ddos.IsSet()
}

// SetDdos sets field value
func (o *IpDetails) SetDdos(v Ddos) {
	o.Ddos.Set(&v)
}

func (o IpDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ip"] = o.Ip
	toSerialize["prefixLength"] = o.PrefixLength
	toSerialize["version"] = o.Version
	toSerialize["nullRouted"] = o.NullRouted
	toSerialize["mainIp"] = o.MainIp
	toSerialize["networkType"] = o.NetworkType
	toSerialize["reverseLookup"] = o.ReverseLookup.Get()
	toSerialize["ddos"] = o.Ddos.Get()
	return toSerialize, nil
}

func (o *IpDetails) UnmarshalJSON(data []byte) (err error) {
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
		"ddos",
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

	varIpDetails := _IpDetails{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIpDetails)

	if err != nil {
		return err
	}

	*o = IpDetails(varIpDetails)

	return err
}

type NullableIpDetails struct {
	value *IpDetails
	isSet bool
}

func (v NullableIpDetails) Get() *IpDetails {
	return v.value
}

func (v *NullableIpDetails) Set(val *IpDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableIpDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableIpDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpDetails(val *IpDetails) *NullableIpDetails {
	return &NullableIpDetails{value: val, isSet: true}
}

func (v NullableIpDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


