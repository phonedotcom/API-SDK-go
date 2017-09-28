# \OauthclientsredirecturisApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountOauthClientsRedirectUri**](OauthclientsredirecturisApi.md#CreateAccountOauthClientsRedirectUri) | **Post** /accounts/{account_id}/oauth/clients/{client_id}/redirect-uris | Create an OAuth Client Redirect URI record.
[**DeleteAccountOauthClientsRedirectUri**](OauthclientsredirecturisApi.md#DeleteAccountOauthClientsRedirectUri) | **Delete** /accounts/{account_id}/oauth/clients/{client_id}/redirect-uris/{uri_id} | Delete an OAuth Client Redirect URI record.
[**GetAccountOauthClientsRedirectUri**](OauthclientsredirecturisApi.md#GetAccountOauthClientsRedirectUri) | **Get** /accounts/{account_id}/oauth/clients/{client_id}/redirect-uris/{uri_id} | Get details of an OAuth Client Redirect URI record.
[**ListAccountOauthClientsRedirectUris**](OauthclientsredirecturisApi.md#ListAccountOauthClientsRedirectUris) | **Get** /accounts/{account_id}/oauth/clients/{client_id}/redirect-uris | Get a list of OAuth Client Redirect URIs for an account.


# **CreateAccountOauthClientsRedirectUri**
> OauthClientRedirectUriFull CreateAccountOauthClientsRedirectUri($accountId, $clientId, $data)

Create an OAuth Client Redirect URI record.

Create an OAuth Client Redirect URI record.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **clientId** | **int32**| Client ID | 
 **data** | [**CreateRedirectUriParams**](CreateRedirectUriParams.md)| Redirect Uri data | [optional] 

### Return type

[**OauthClientRedirectUriFull**](OauthClientRedirectUriFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountOauthClientsRedirectUri**
> DeleteEntry DeleteAccountOauthClientsRedirectUri($accountId, $clientId, $uriId)

Delete an OAuth Client Redirect URI record.

Delete an OAuth Client Redirect URI record.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **clientId** | **int32**| Client ID | 
 **uriId** | **int32**| Redirect URI ID | 

### Return type

[**DeleteEntry**](DeleteEntry.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountOauthClientsRedirectUri**
> OauthClientRedirectUriFull GetAccountOauthClientsRedirectUri($accountId, $clientId, $uriId)

Get details of an OAuth Client Redirect URI record.

Get details of an OAuth Client Redirect URI record.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **clientId** | **int32**| Client ID | 
 **uriId** | **int32**| Redirect URI ID | 

### Return type

[**OauthClientRedirectUriFull**](OauthClientRedirectUriFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountOauthClientsRedirectUris**
> ListOauthClientsRedirectUris ListAccountOauthClientsRedirectUris($accountId, $clientId, $filtersId, $sortId, $limit, $offset, $fields)

Get a list of OAuth Client Redirect URIs for an account.

Get a list of OAuth Client Redirect URIs for an account.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **clientId** | **int32**| Client ID | 
 **filtersId** | [**[]string**](string.md)| ID filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListOauthClientsRedirectUris**](ListOauthClientsRedirectUris.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

