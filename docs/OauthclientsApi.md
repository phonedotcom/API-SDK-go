# \OauthclientsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccountOauthClient**](OauthclientsApi.md#DeleteAccountOauthClient) | **Delete** /accounts/{account_id}/oauth/clients/{client_id} | Delete an individual OAuth client.
[**GetAccountOauthClient**](OauthclientsApi.md#GetAccountOauthClient) | **Get** /accounts/{account_id}/oauth/clients/{client_id} | Show details of an individual OAuth client.
[**ListAccountOauthClients**](OauthclientsApi.md#ListAccountOauthClients) | **Get** /accounts/{account_id}/oauth/clients | Get a list of OAuth clients for an account.


# **DeleteAccountOauthClient**
> DeleteEntry DeleteAccountOauthClient($accountId, $clientId)

Delete an individual OAuth client.

Delete an individual OAuth client. See Account OAuth Clients for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **clientId** | **int32**| Client ID | 

### Return type

[**DeleteEntry**](DeleteEntry.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountOauthClient**
> OauthClientFull GetAccountOauthClient($accountId, $clientId)

Show details of an individual OAuth client.

Show details of an individual OAuth client. See Account OAuth Clients for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **clientId** | **int32**| Client ID | 

### Return type

[**OauthClientFull**](OauthClientFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountOauthClients**
> ListOauthClients ListAccountOauthClients($accountId, $filtersId, $sortId, $limit, $offset, $fields)

Get a list of OAuth clients for an account.

Get a list of OAuth clients for an account. See Account OAuth Clients for more info on the properties.


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

[**ListOauthClients**](ListOauthClients.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

