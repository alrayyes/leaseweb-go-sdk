# Go API client for invoice

> The base URL for this API is: **https://api.leaseweb.com/invoices/v1/_**

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Generator version: 7.10.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import invoice "github.com/leaseweb/leaseweb-go-sdk/invoice"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `invoice.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), invoice.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `invoice.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), invoice.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `invoice.ContextOperationServerIndices` and `invoice.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), invoice.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), invoice.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.leaseweb.com/invoices/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*InvoiceAPI* | [**GetInvoice**](docs/InvoiceAPI.md#getinvoice) | **Get** /invoices/{invoiceId} | Inspect an invoice
*InvoiceAPI* | [**GetInvoiceList**](docs/InvoiceAPI.md#getinvoicelist) | **Get** /invoices | List invoices
*InvoiceAPI* | [**GetInvoicePdf**](docs/InvoiceAPI.md#getinvoicepdf) | **Get** /invoices/{invoiceId}/pdf | Get invoice PDF
*InvoiceAPI* | [**GetProforma**](docs/InvoiceAPI.md#getproforma) | **Get** /invoices/proforma | Pro Forma


## Documentation For Models

 - [ContractItem](docs/ContractItem.md)
 - [Credit](docs/Credit.md)
 - [ErrorResult](docs/ErrorResult.md)
 - [GetInvoiceListResult](docs/GetInvoiceListResult.md)
 - [GetInvoiceResult](docs/GetInvoiceResult.md)
 - [GetProformaResult](docs/GetProformaResult.md)
 - [Invoice](docs/Invoice.md)
 - [LineItem](docs/LineItem.md)
 - [Metadata](docs/Metadata.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BearerAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), invoice.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```

### X-LSW-Auth

- **Type**: API key
- **API key parameter name**: X-LSW-Auth
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-LSW-Auth and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		invoice.ContextAPIKeys,
		map[string]invoice.APIKey{
			"X-LSW-Auth": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



