/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
)

// checks if the Model400ErrorResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model400ErrorResult{}

// Model400ErrorResult struct for Model400ErrorResult
type Model400ErrorResult struct {
	// The correlation ID of the current request.
	CorrelationId *string `json:"correlationId,omitempty"`
	// The error code.
	ErrorCode *string `json:"errorCode,omitempty"`
	// A human friendly description of the error.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	ErrorDetails *map[string][]string `json:"errorDetails,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Model400ErrorResult Model400ErrorResult

// NewModel400ErrorResult instantiates a new Model400ErrorResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel400ErrorResult() *Model400ErrorResult {
	this := Model400ErrorResult{}
	return &this
}

// NewModel400ErrorResultWithDefaults instantiates a new Model400ErrorResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel400ErrorResultWithDefaults() *Model400ErrorResult {
	this := Model400ErrorResult{}
	return &this
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *Model400ErrorResult) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model400ErrorResult) GetCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *Model400ErrorResult) HasCorrelationId() bool {
	if o != nil && !IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *Model400ErrorResult) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *Model400ErrorResult) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model400ErrorResult) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *Model400ErrorResult) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *Model400ErrorResult) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *Model400ErrorResult) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model400ErrorResult) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *Model400ErrorResult) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *Model400ErrorResult) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *Model400ErrorResult) GetErrorDetails() map[string][]string {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret map[string][]string
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model400ErrorResult) GetErrorDetailsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *Model400ErrorResult) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given map[string][]string and assigns it to the ErrorDetails field.
func (o *Model400ErrorResult) SetErrorDetails(v map[string][]string) {
	o.ErrorDetails = &v
}

func (o Model400ErrorResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model400ErrorResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CorrelationId) {
		toSerialize["correlationId"] = o.CorrelationId
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Model400ErrorResult) UnmarshalJSON(data []byte) (err error) {
	varModel400ErrorResult := _Model400ErrorResult{}

	err = json.Unmarshal(data, &varModel400ErrorResult)

	if err != nil {
		return err
	}

	*o = Model400ErrorResult(varModel400ErrorResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "correlationId")
		delete(additionalProperties, "errorCode")
		delete(additionalProperties, "errorMessage")
		delete(additionalProperties, "errorDetails")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModel400ErrorResult struct {
	value *Model400ErrorResult
	isSet bool
}

func (v NullableModel400ErrorResult) Get() *Model400ErrorResult {
	return v.value
}

func (v *NullableModel400ErrorResult) Set(val *Model400ErrorResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModel400ErrorResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModel400ErrorResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel400ErrorResult(val *Model400ErrorResult) *NullableModel400ErrorResult {
	return &NullableModel400ErrorResult{value: val, isSet: true}
}

func (v NullableModel400ErrorResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel400ErrorResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


