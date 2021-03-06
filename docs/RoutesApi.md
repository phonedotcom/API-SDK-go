# \RoutesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoute**](RoutesApi.md#CreateRoute) | **Post** /accounts/{account_id}/routes | Add a new route to the account.
[**DeleteAccountRoute**](RoutesApi.md#DeleteAccountRoute) | **Delete** /accounts/{account_id}/routes/{route_id} | Delete a route from the account.
[**GetAccountRoute**](RoutesApi.md#GetAccountRoute) | **Get** /accounts/{account_id}/routes/{route_id} | Show details of an individual route.
[**ListAccountRoutes**](RoutesApi.md#ListAccountRoutes) | **Get** /accounts/{account_id}/routes | Get a list of routes for an account.
[**ReplaceAccountRoute**](RoutesApi.md#ReplaceAccountRoute) | **Put** /accounts/{account_id}/routes/{route_id} | Update the information of a route.


# **CreateRoute**
> RouteFull CreateRoute($accountId, $data)

Add a new route to the account.

Add a new route to the account. See Intro to Routes for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **data** | [**CreateRouteParams**](CreateRouteParams.md)| Route data | [optional] 

### Return type

[**RouteFull**](RouteFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountRoute**
> DeleteEntry DeleteAccountRoute($accountId, $routeId)

Delete a route from the account.

Delete a route from the account. See Intro to Routes for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **routeId** | **int32**| Route ID | 

### Return type

[**DeleteEntry**](DeleteEntry.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountRoute**
> RouteFull GetAccountRoute($accountId, $routeId)

Show details of an individual route.

Show details of an individual route. See Intro to Routes for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **routeId** | **int32**| Route ID | 

### Return type

[**RouteFull**](RouteFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountRoutes**
> ListRoutes ListAccountRoutes($accountId, $filtersId, $filtersName, $sortId, $sortName, $limit, $offset, $fields)

Get a list of routes for an account.

Get a list of routes for an account. See Intro to Routes for more info on the properties.


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

[**ListRoutes**](ListRoutes.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAccountRoute**
> RouteFull ReplaceAccountRoute($accountId, $routeId, $data)

Update the information of a route.

Update the information of a route. See Intro to Routes for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **routeId** | **int32**| Route ID | 
 **data** | [**CreateRouteParams**](CreateRouteParams.md)| Route data | [optional] 

### Return type

[**RouteFull**](RouteFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

