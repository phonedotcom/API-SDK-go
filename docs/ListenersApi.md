# \ListenersApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountListener**](ListenersApi.md#CreateAccountListener) | **Post** /accounts/{account_id}/listeners | Add a listener object to your account that can be used to subscribe an event.
[**DeleteAccountListener**](ListenersApi.md#DeleteAccountListener) | **Delete** /accounts/{account_id}/listeners/{listener_id} | Delete an individual event listener.
[**GetAccountListener**](ListenersApi.md#GetAccountListener) | **Get** /accounts/{account_id}/listeners/{listener_id} | Show details of an individual listener.
[**ListAccountListeners**](ListenersApi.md#ListAccountListeners) | **Get** /accounts/{account_id}/listeners | Get a list of listeners for an account.
[**ReplaceAccountListener**](ListenersApi.md#ReplaceAccountListener) | **Put** /accounts/{account_id}/listeners/{listener_id} | Update the settings of an individual event listener.


# **CreateAccountListener**
> ListenerFull CreateAccountListener($accountId, $data)

Add a listener object to your account that can be used to subscribe an event.

Add a listener object to your account that can be used to subscribe an event. See Account Listeners for more info on the properties. Note: This API is for users with Account Owner scope access token. Users with Extension User scope token should invoke the Extension level Post Listener API with the following definition: POST https://api.phone.com/v4/accounts/:account_id/extensions/:extension_id/listeners


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **data** | [**CreateListenerParams**](CreateListenerParams.md)| Account Listeners Data | [optional] 

### Return type

[**ListenerFull**](ListenerFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountListener**
> DeleteEntry DeleteAccountListener($accountId, $listenerId)

Delete an individual event listener.

Delete an individual event listener. See Account Listeners for more info on the properties. Note: This API is for users with Account Owner scope access token. Users with Extension User scope token should invoke the Extension level Delete Listener API with the following definition: DELETE https://api.phone.com/v4/accounts/:account_id/extensions/:extension_id/listeners/:listener_id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **listenerId** | **int32**| Listener ID | 

### Return type

[**DeleteEntry**](DeleteEntry.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountListener**
> ListenerFull GetAccountListener($accountId, $listenerId)

Show details of an individual listener.

Show details of an individual event listener. See Account Listeners for more info on the properties. Note: This API is for users with Account Owner scope access token. Users with Extension User scope token should invoke the Extension level Get Listener API with the following definition: GET https://api.phone.com/v4/accounts/:account_id/extensions/:extension_id/listeners/:listener_id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **listenerId** | **int32**| Listener ID | 

### Return type

[**ListenerFull**](ListenerFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountListeners**
> ListListeners ListAccountListeners($accountId, $filtersId, $sortId, $limit, $offset, $fields)

Get a list of listeners for an account.

Get a list of listeners for an account. See Account Listeners for more info on the properties. Note: This API is for users with Account Owner scope access token. Users with Extension User scope token should invoke the Extension level List Listeners API with the following definition: GET https://api.phone.com/v4/accounts/:account_id/extensions/:extension_id/listeners


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

[**ListListeners**](ListListeners.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAccountListener**
> ListenerFull ReplaceAccountListener($accountId, $listenerId, $data)

Update the settings of an individual event listener.

Update the settings of an individual event listener. See Event Listeners for more detail. Note: This API is for users with Account Owner scope access token. Users with Extension User scope token should invoke the Extension level Replace Listener API with the following definition: PUT https://api.phone.com/v4/accounts/:account_id/extensions/:extension_id/listeners/:listener_id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **listenerId** | **int32**| Listener ID | 
 **data** | [**CreateListenerParams**](CreateListenerParams.md)| Account Listeners Data | [optional] 

### Return type

[**ListenerFull**](ListenerFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

