# \VoicemailApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountVoicemail**](VoicemailApi.md#GetAccountVoicemail) | **Get** /accounts/{account_id}/voicemail/{voicemail_id} | This service shows the details of an individual voicemail.
[**ListAccountVoicemail**](VoicemailApi.md#ListAccountVoicemail) | **Get** /accounts/{account_id}/voicemail | Get a list of voicemail messages for an account.
[**PatchAccountVoicemail**](VoicemailApi.md#PatchAccountVoicemail) | **Patch** /accounts/{account_id}/voicemail/{voicemail_id} | Update the is_new parameter in a voicemail record.


# **GetAccountVoicemail**
> VoicemailFull GetAccountVoicemail($accountId, $voicemailId)

This service shows the details of an individual voicemail.

This service shows the details of an individual voicemail. See Intro to Voicemail for more info on the properties. Note: This API is for users with Account Owner scope access token. Users with Extension User scope token should invoke the Get Voicemail API with the following definition: GET https://api.phone.com/v4/accounts/:account_id/extensions/:extension_id/voicemail/:voicemail_id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **voicemailId** | **string**| Voicemail ID | 

### Return type

[**VoicemailFull**](VoicemailFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountVoicemail**
> ListVoicemail ListAccountVoicemail($accountId, $filtersId, $filtersFrom, $filtersTo, $filtersIsNew, $filtersCreatedAt, $filtersExtension, $sortId, $sortCreatedAt, $limit, $offset, $fields)

Get a list of voicemail messages for an account.

Get a list of voicemail messages for an account. See Intro to Voicemail for more info on the properties. Note: This API is for users with Account Owner scope access token. Users with Extension User scope token should invoke the List Voicemail API with the following definition: GET https://api.phone.com/v4/accounts/:account_id/extensions/:extension_id/voicemail


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **filtersId** | [**[]string**](string.md)| ID filter | [optional] 
 **filtersFrom** | **string**| Caller ID filter | [optional] 
 **filtersTo** | **string**| Callee ID filter, the E.164 phone number to send the SMS TO. Note you must encode the + as %2B | [optional] 
 **filtersIsNew** | **bool**|  | [optional] 
 **filtersCreatedAt** | **string**| Date string representing the UTC time that sms was created | [optional] 
 **filtersExtension** | [**[]string**](string.md)| Extension filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortCreatedAt** | **string**| Sort by the UTC time that voicemail was created | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListVoicemail**](ListVoicemail.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchAccountVoicemail**
> VoicemailFull PatchAccountVoicemail($accountId, $voicemailId, $data)

Update the is_new parameter in a voicemail record.

Update the is_new parameter in a voicemail record. See Account Voicemail for more info on the properties. Note: This API is for users with Account Owner scope access token. Users with Extension User scope token should invoke the Patch Voicemail API with the following definition: PATCH https://api.phone.com/v4/accounts/:account_id/extensions/:extension_id/voicemail/:voicemail_id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **voicemailId** | **string**| Voicemail ID | 
 **data** | [**PatchVoicemailParams**](PatchVoicemailParams.md)| Payment data | [optional] 

### Return type

[**VoicemailFull**](VoicemailFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

