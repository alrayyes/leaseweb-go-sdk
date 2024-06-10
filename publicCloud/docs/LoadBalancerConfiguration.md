# LoadBalancerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StickySession** | Pointer to [**NullableStickySession**](StickySession.md) |  | [optional] 
**Balance** | Pointer to **string** | Algorithm to be used for load balancer | [optional] 
**HealthCheck** | Pointer to [**NullableHealthCheck**](HealthCheck.md) |  | [optional] 
**XForwardedFor** | Pointer to **bool** | Is xForwardedFor header enabled or not | [optional] 
**IdleTimeOut** | Pointer to **int32** | Time to close the connection if load balancer is idle | [optional] 
**TargetPort** | Pointer to **int32** | Port on which the backend (target) servers are listening to handle incoming requests | [optional] 

## Methods

### NewLoadBalancerConfiguration

`func NewLoadBalancerConfiguration() *LoadBalancerConfiguration`

NewLoadBalancerConfiguration instantiates a new LoadBalancerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerConfigurationWithDefaults

`func NewLoadBalancerConfigurationWithDefaults() *LoadBalancerConfiguration`

NewLoadBalancerConfigurationWithDefaults instantiates a new LoadBalancerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStickySession

`func (o *LoadBalancerConfiguration) GetStickySession() StickySession`

GetStickySession returns the StickySession field if non-nil, zero value otherwise.

### GetStickySessionOk

`func (o *LoadBalancerConfiguration) GetStickySessionOk() (*StickySession, bool)`

GetStickySessionOk returns a tuple with the StickySession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickySession

`func (o *LoadBalancerConfiguration) SetStickySession(v StickySession)`

SetStickySession sets StickySession field to given value.

### HasStickySession

`func (o *LoadBalancerConfiguration) HasStickySession() bool`

HasStickySession returns a boolean if a field has been set.

### SetStickySessionNil

`func (o *LoadBalancerConfiguration) SetStickySessionNil(b bool)`

 SetStickySessionNil sets the value for StickySession to be an explicit nil

### UnsetStickySession
`func (o *LoadBalancerConfiguration) UnsetStickySession()`

UnsetStickySession ensures that no value is present for StickySession, not even an explicit nil
### GetBalance

`func (o *LoadBalancerConfiguration) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *LoadBalancerConfiguration) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *LoadBalancerConfiguration) SetBalance(v string)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *LoadBalancerConfiguration) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetHealthCheck

`func (o *LoadBalancerConfiguration) GetHealthCheck() HealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *LoadBalancerConfiguration) GetHealthCheckOk() (*HealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *LoadBalancerConfiguration) SetHealthCheck(v HealthCheck)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *LoadBalancerConfiguration) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### SetHealthCheckNil

`func (o *LoadBalancerConfiguration) SetHealthCheckNil(b bool)`

 SetHealthCheckNil sets the value for HealthCheck to be an explicit nil

### UnsetHealthCheck
`func (o *LoadBalancerConfiguration) UnsetHealthCheck()`

UnsetHealthCheck ensures that no value is present for HealthCheck, not even an explicit nil
### GetXForwardedFor

`func (o *LoadBalancerConfiguration) GetXForwardedFor() bool`

GetXForwardedFor returns the XForwardedFor field if non-nil, zero value otherwise.

### GetXForwardedForOk

`func (o *LoadBalancerConfiguration) GetXForwardedForOk() (*bool, bool)`

GetXForwardedForOk returns a tuple with the XForwardedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXForwardedFor

`func (o *LoadBalancerConfiguration) SetXForwardedFor(v bool)`

SetXForwardedFor sets XForwardedFor field to given value.

### HasXForwardedFor

`func (o *LoadBalancerConfiguration) HasXForwardedFor() bool`

HasXForwardedFor returns a boolean if a field has been set.

### GetIdleTimeOut

`func (o *LoadBalancerConfiguration) GetIdleTimeOut() int32`

GetIdleTimeOut returns the IdleTimeOut field if non-nil, zero value otherwise.

### GetIdleTimeOutOk

`func (o *LoadBalancerConfiguration) GetIdleTimeOutOk() (*int32, bool)`

GetIdleTimeOutOk returns a tuple with the IdleTimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeOut

`func (o *LoadBalancerConfiguration) SetIdleTimeOut(v int32)`

SetIdleTimeOut sets IdleTimeOut field to given value.

### HasIdleTimeOut

`func (o *LoadBalancerConfiguration) HasIdleTimeOut() bool`

HasIdleTimeOut returns a boolean if a field has been set.

### GetTargetPort

`func (o *LoadBalancerConfiguration) GetTargetPort() int32`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *LoadBalancerConfiguration) GetTargetPortOk() (*int32, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *LoadBalancerConfiguration) SetTargetPort(v int32)`

SetTargetPort sets TargetPort field to given value.

### HasTargetPort

`func (o *LoadBalancerConfiguration) HasTargetPort() bool`

HasTargetPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

