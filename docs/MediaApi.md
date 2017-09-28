# \MediaApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountMediaFiles**](MediaApi.md#CreateAccountMediaFiles) | **Post** /accounts/{account_id}/media/files | Add a media object to your account that can be used as a greeting or hold music. Users may create a media by using the built-in Text-to-speech (TTS) facility or upload a file of their choice. (Note: The maximum size for media files or JSON objects included with a POST or PUT request is 10 MB)
[**CreateAccountMediaTts**](MediaApi.md#CreateAccountMediaTts) | **Post** /accounts/{account_id}/media/tts | Add a media object to your account that can be used as a greeting or hold music. Users may create a media by using the built-in Text-to-speech (TTS) facility or upload a file of their choice. (Note: The maximum size for media files or JSON objects included with a POST or PUT request is 10 MB)
[**DeleteAccountMedia**](MediaApi.md#DeleteAccountMedia) | **Delete** /accounts/{account_id}/media/{media_id} | Delete an individual media record
[**GetAccountMedia**](MediaApi.md#GetAccountMedia) | **Get** /accounts/{account_id}/media/{media_id} | Show details of an individual media recording (Greeting or Hold Music)
[**ListAccountMedia**](MediaApi.md#ListAccountMedia) | **Get** /accounts/{account_id}/media | Get a list of media recordings for an account.
[**ReplaceAccountMediaFiles**](MediaApi.md#ReplaceAccountMediaFiles) | **Put** /accounts/{account_id}/media/files/{media_id} | Update a media object to your account. Note: The maximum size for media files or JSON objects included with a POST or PUT request is 10 MB.
[**ReplaceAccountMediaTts**](MediaApi.md#ReplaceAccountMediaTts) | **Put** /accounts/{account_id}/media/tts/{media_id} | Update a media object to your account.


# **CreateAccountMediaFiles**
> MediaFull CreateAccountMediaFiles($accountId, $json, $file)

Add a media object to your account that can be used as a greeting or hold music. Users may create a media by using the built-in Text-to-speech (TTS) facility or upload a file of their choice. (Note: The maximum size for media files or JSON objects included with a POST or PUT request is 10 MB)

See Account Media for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **json** | **string**| Media extra parameters | [optional] 
 **file** | ***os.File**| Media file | [optional] 

### Return type

[**MediaFull**](MediaFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccountMediaTts**
> MediaFull CreateAccountMediaTts($accountId, $data)

Add a media object to your account that can be used as a greeting or hold music. Users may create a media by using the built-in Text-to-speech (TTS) facility or upload a file of their choice. (Note: The maximum size for media files or JSON objects included with a POST or PUT request is 10 MB)

See Account Media for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **data** | [**CreateMediaParams**](CreateMediaParams.md)| Media data | [optional] 

### Return type

[**MediaFull**](MediaFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountMedia**
> DeleteEntry DeleteAccountMedia($accountId, $mediaId)

Delete an individual media record

See Account Media for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **mediaId** | **int32**| Media ID | 

### Return type

[**DeleteEntry**](DeleteEntry.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountMedia**
> MediaFull GetAccountMedia($accountId, $mediaId)

Show details of an individual media recording (Greeting or Hold Music)

Get individual media recording


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **mediaId** | **int32**| Media ID | 

### Return type

[**MediaFull**](MediaFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountMedia**
> ListMedia ListAccountMedia($accountId, $filtersId, $filtersName, $sortId, $sortName, $limit, $offset, $fields)

Get a list of media recordings for an account.

Get a list of media recordings for an account. See Account Media for more info on the properties. Note: This API is for users with Account Owner scope access token. Users with Extension User scope token should invoke the Extension level List Media API with the following definition: GET https://api.phone.com/v4/accounts/:account_id/extensions/:extension_id/media


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

[**ListMedia**](ListMedia.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAccountMediaFiles**
> MediaFull ReplaceAccountMediaFiles($accountId, $mediaId, $json, $file)

Update a media object to your account. Note: The maximum size for media files or JSON objects included with a POST or PUT request is 10 MB.

See Account Media for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **mediaId** | **int32**| Media ID | 
 **json** | **string**| Media extra parameters | [optional] 
 **file** | ***os.File**| Media file | [optional] 

### Return type

[**MediaFull**](MediaFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAccountMediaTts**
> MediaFull ReplaceAccountMediaTts($accountId, $mediaId, $data)

Update a media object to your account.

Update a media object to your account. Note: The maximum size for media files or JSON objects included with a POST or PUT request is 10 MB. See Account Media for more info on the properties. Note: This API is for users with Account Owner scope access token. Users with Extension User scope token should invoke the Extension level Replace Media API with the following definition: PUT https://api.phone.com/v4/accounts/:account_id/extensions/:extension_id/media/:media_id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **mediaId** | **int32**| Media ID | 
 **data** | [**CreateMediaParams**](CreateMediaParams.md)| Media data | [optional] 

### Return type

[**MediaFull**](MediaFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

