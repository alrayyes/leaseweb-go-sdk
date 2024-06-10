/*
Invoices

> The base URL for this API is: **https://api.leaseweb.com/invoices/v1/_**

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package invoice

import (
	"encoding/json"
)

// checks if the GetProformaResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProformaResult{}

// GetProformaResult struct for GetProformaResult
type GetProformaResult struct {
	Metadata *Metadata `json:"_metadata,omitempty"`
	ContractItems []ContractItem `json:"contractItems,omitempty"`
	// The currency of the invoice. Based on ISO 4217
	Currency *string `json:"currency,omitempty"`
	// The date of the next invoice on which this proforma is based on.
	ProformaDate *string `json:"proformaDate,omitempty"`
	// Total amount without vat which will be invoiced the upcoming month.
	SubTotal *float32 `json:"subTotal,omitempty"`
	// Total amount which will be invoiced the upcoming month.
	Total *float32 `json:"total,omitempty"`
	// The total amount of vat.
	VatAmount *float32 `json:"vatAmount,omitempty"`
}

// NewGetProformaResult instantiates a new GetProformaResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProformaResult() *GetProformaResult {
	this := GetProformaResult{}
	return &this
}

// NewGetProformaResultWithDefaults instantiates a new GetProformaResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProformaResultWithDefaults() *GetProformaResult {
	this := GetProformaResult{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetProformaResult) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProformaResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetProformaResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *GetProformaResult) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetContractItems returns the ContractItems field value if set, zero value otherwise.
func (o *GetProformaResult) GetContractItems() []ContractItem {
	if o == nil || IsNil(o.ContractItems) {
		var ret []ContractItem
		return ret
	}
	return o.ContractItems
}

// GetContractItemsOk returns a tuple with the ContractItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProformaResult) GetContractItemsOk() ([]ContractItem, bool) {
	if o == nil || IsNil(o.ContractItems) {
		return nil, false
	}
	return o.ContractItems, true
}

// HasContractItems returns a boolean if a field has been set.
func (o *GetProformaResult) HasContractItems() bool {
	if o != nil && !IsNil(o.ContractItems) {
		return true
	}

	return false
}

// SetContractItems gets a reference to the given []ContractItem and assigns it to the ContractItems field.
func (o *GetProformaResult) SetContractItems(v []ContractItem) {
	o.ContractItems = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GetProformaResult) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProformaResult) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GetProformaResult) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *GetProformaResult) SetCurrency(v string) {
	o.Currency = &v
}

// GetProformaDate returns the ProformaDate field value if set, zero value otherwise.
func (o *GetProformaResult) GetProformaDate() string {
	if o == nil || IsNil(o.ProformaDate) {
		var ret string
		return ret
	}
	return *o.ProformaDate
}

// GetProformaDateOk returns a tuple with the ProformaDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProformaResult) GetProformaDateOk() (*string, bool) {
	if o == nil || IsNil(o.ProformaDate) {
		return nil, false
	}
	return o.ProformaDate, true
}

// HasProformaDate returns a boolean if a field has been set.
func (o *GetProformaResult) HasProformaDate() bool {
	if o != nil && !IsNil(o.ProformaDate) {
		return true
	}

	return false
}

// SetProformaDate gets a reference to the given string and assigns it to the ProformaDate field.
func (o *GetProformaResult) SetProformaDate(v string) {
	o.ProformaDate = &v
}

// GetSubTotal returns the SubTotal field value if set, zero value otherwise.
func (o *GetProformaResult) GetSubTotal() float32 {
	if o == nil || IsNil(o.SubTotal) {
		var ret float32
		return ret
	}
	return *o.SubTotal
}

// GetSubTotalOk returns a tuple with the SubTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProformaResult) GetSubTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.SubTotal) {
		return nil, false
	}
	return o.SubTotal, true
}

// HasSubTotal returns a boolean if a field has been set.
func (o *GetProformaResult) HasSubTotal() bool {
	if o != nil && !IsNil(o.SubTotal) {
		return true
	}

	return false
}

// SetSubTotal gets a reference to the given float32 and assigns it to the SubTotal field.
func (o *GetProformaResult) SetSubTotal(v float32) {
	o.SubTotal = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetProformaResult) GetTotal() float32 {
	if o == nil || IsNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProformaResult) GetTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetProformaResult) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *GetProformaResult) SetTotal(v float32) {
	o.Total = &v
}

// GetVatAmount returns the VatAmount field value if set, zero value otherwise.
func (o *GetProformaResult) GetVatAmount() float32 {
	if o == nil || IsNil(o.VatAmount) {
		var ret float32
		return ret
	}
	return *o.VatAmount
}

// GetVatAmountOk returns a tuple with the VatAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProformaResult) GetVatAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.VatAmount) {
		return nil, false
	}
	return o.VatAmount, true
}

// HasVatAmount returns a boolean if a field has been set.
func (o *GetProformaResult) HasVatAmount() bool {
	if o != nil && !IsNil(o.VatAmount) {
		return true
	}

	return false
}

// SetVatAmount gets a reference to the given float32 and assigns it to the VatAmount field.
func (o *GetProformaResult) SetVatAmount(v float32) {
	o.VatAmount = &v
}

func (o GetProformaResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProformaResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}
	if !IsNil(o.ContractItems) {
		toSerialize["contractItems"] = o.ContractItems
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.ProformaDate) {
		toSerialize["proformaDate"] = o.ProformaDate
	}
	if !IsNil(o.SubTotal) {
		toSerialize["subTotal"] = o.SubTotal
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.VatAmount) {
		toSerialize["vatAmount"] = o.VatAmount
	}
	return toSerialize, nil
}

type NullableGetProformaResult struct {
	value *GetProformaResult
	isSet bool
}

func (v NullableGetProformaResult) Get() *GetProformaResult {
	return v.value
}

func (v *NullableGetProformaResult) Set(val *GetProformaResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProformaResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProformaResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProformaResult(val *GetProformaResult) *NullableGetProformaResult {
	return &NullableGetProformaResult{value: val, isSet: true}
}

func (v NullableGetProformaResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProformaResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

