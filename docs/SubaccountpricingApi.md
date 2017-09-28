# \SubaccountpricingApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountSubaccountPricing**](SubaccountpricingApi.md#CreateAccountSubaccountPricing) | **Post** /accounts/{account_id}/subaccounts/{subaccount_id}/pricing | Add a pricing plan to a subaccount.
[**DeleteAccountSubaccountPricing**](SubaccountpricingApi.md#DeleteAccountSubaccountPricing) | **Delete** /accounts/{account_id}/subaccounts/{subaccount_id}/pricing/{pricing_id} | Delete a pricing plan from a subaccount.
[**GetAccountSubaccountPricing**](SubaccountpricingApi.md#GetAccountSubaccountPricing) | **Get** /accounts/{account_id}/subaccounts/{subaccount_id}/pricing/{pricing_id} | Get the details of a pricing plan for a subaccount.
[**ListAccountSubaccountPricing**](SubaccountpricingApi.md#ListAccountSubaccountPricing) | **Get** /accounts/{account_id}/subaccounts/{subaccount_id}/pricing | Get a list of pricing plans for a subaccount.


# **CreateAccountSubaccountPricing**
> PricingFull CreateAccountSubaccountPricing($accountId, $subaccountId, $data)

Add a pricing plan to a subaccount.

Add a pricing plan to a subaccount. See Account Subaccount Pricing for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **subaccountId** | **int32**| Subaccount ID | 
 **data** | [**CreatePricingParams**](CreatePricingParams.md)| Subaccount pricing data | 

### Return type

[**PricingFull**](PricingFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountSubaccountPricing**
> DeleteEntry DeleteAccountSubaccountPricing($accountId, $subaccountId, $pricingId)

Delete a pricing plan from a subaccount.

Delete a pricing plan from a subaccount. See Account Subaccount Pricing for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **subaccountId** | **int32**| Subaccount ID | 
 **pricingId** | **int32**| Pricing Object ID | 

### Return type

[**DeleteEntry**](DeleteEntry.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountSubaccountPricing**
> PricingFull GetAccountSubaccountPricing($accountId, $subaccountId, $pricingId)

Get the details of a pricing plan for a subaccount.

Get the details of a pricing plan for a subaccount. See Account Subaccount Pricing for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **subaccountId** | **int32**| Subaccount ID | 
 **pricingId** | **int32**| Pricing Object ID | 

### Return type

[**PricingFull**](PricingFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountSubaccountPricing**
> ListPricings ListAccountSubaccountPricing($accountId, $subaccountId, $filtersId, $sortId, $limit, $offset, $fields)

Get a list of pricing plans for a subaccount.

Get a list of pricing plans for a subaccount. See Account Subaccount Pricing for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **subaccountId** | **int32**| Subaccount ID | 
 **filtersId** | [**[]string**](string.md)| ID filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListPricings**](ListPricings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

