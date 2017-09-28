# \PaymentmethodsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountPaymentMethod**](PaymentmethodsApi.md#CreateAccountPaymentMethod) | **Post** /accounts/{account_id}/payment-methods | Create an individual payment method.
[**DeleteAccountPaymentMethod**](PaymentmethodsApi.md#DeleteAccountPaymentMethod) | **Delete** /accounts/{account_id}/payment-methods/{pm_id} | Delete an individual payment method.
[**GetAccountPaymentMethod**](PaymentmethodsApi.md#GetAccountPaymentMethod) | **Get** /accounts/{account_id}/payment-methods/{pm_id} | Show details of an individual payment method.
[**ListAccountPaymentMethods**](PaymentmethodsApi.md#ListAccountPaymentMethods) | **Get** /accounts/{account_id}/payment-methods | Get a list of payment methods for an account.
[**PatchAccountPaymentMethod**](PaymentmethodsApi.md#PatchAccountPaymentMethod) | **Patch** /accounts/{account_id}/payment-methods/{pm_id} | Replace the status of an individual payment method.


# **CreateAccountPaymentMethod**
> PaymentFull CreateAccountPaymentMethod($accountId, $data)

Create an individual payment method.

Create an individual payment method. See Account Payment Methods for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **data** | [**CreatePaymentParams**](CreatePaymentParams.md)| Payment data | 

### Return type

[**PaymentFull**](PaymentFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountPaymentMethod**
> DeleteEntry DeleteAccountPaymentMethod($accountId, $pmId)

Delete an individual payment method.

Delete an individual payment method. See Account Payment Methods for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **pmId** | **int32**| Payment Method ID | 

### Return type

[**DeleteEntry**](DeleteEntry.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountPaymentMethod**
> PaymentFull GetAccountPaymentMethod($accountId, $pmId)

Show details of an individual payment method.

Show details of an individual payment method. See Account Payment Methods for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **pmId** | **int32**| Payment Method ID | 

### Return type

[**PaymentFull**](PaymentFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountPaymentMethods**
> ListPaymentMethods ListAccountPaymentMethods($accountId, $filtersId, $sortId, $limit, $offset, $fields)

Get a list of payment methods for an account.

Get a list of payment methods for an account. See Account Payment Methods for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **filtersId** | [**[]string**](string.md)| ID filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListPaymentMethods**](ListPaymentMethods.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchAccountPaymentMethod**
> PaymentFull PatchAccountPaymentMethod($accountId, $pmId, $data)

Replace the status of an individual payment method.

Replace the status of an individual payment method. See Account Payment Methods for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **pmId** | **int32**| Payment Method ID | 
 **data** | [**PatchPaymentParams**](PatchPaymentParams.md)| Payment data | [optional] 

### Return type

[**PaymentFull**](PaymentFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

