# \ApplicationsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountApplication**](ApplicationsApi.md#GetAccountApplication) | **Get** /accounts/{account_id}/applications/{application_id} | Show details of an individual Application on a given account.
[**ListAccountApplications**](ApplicationsApi.md#ListAccountApplications) | **Get** /accounts/{account_id}/applications | This service lists the Applications on a given account


# **GetAccountApplication**
> ApplicationFull GetAccountApplication($accountId, $applicationId)

Show details of an individual Application on a given account.

Show details of an individual Application on a given account.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **applicationId** | **int32**| Application ID | 

### Return type

[**ApplicationFull**](ApplicationFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountApplications**
> ListApplications ListAccountApplications($accountId, $filtersId, $filtersName, $sortId, $sortName, $limit, $offset, $fields)

This service lists the Applications on a given account

Show details of an individual Application on a given account.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **filtersId** | [**[]string**](string.md)| ID filter | [optional] 
 **filtersName** | [**[]string**](string.md)| Name filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortName** | **string**| Name sorting | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListApplications**](ListApplications.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

