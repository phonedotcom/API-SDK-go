# \SmsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountSms**](SmsApi.md#CreateAccountSms) | **Post** /accounts/{account_id}/sms | Send a SMS to one or a group of recipients.
[**GetAccountSms**](SmsApi.md#GetAccountSms) | **Get** /accounts/{account_id}/sms/{sms_id} | This service shows the details of an individual SMS.
[**ListAccountSms**](SmsApi.md#ListAccountSms) | **Get** /accounts/{account_id}/sms | Get a list of SMS messages for an account.
[**PatchAccountSms**](SmsApi.md#PatchAccountSms) | **Patch** /accounts/{account_id}/sms/{sms_id} | Update the is_new parameter in a sms record.


# **CreateAccountSms**
> SmsFull CreateAccountSms($accountId, $data)

Send a SMS to one or a group of recipients.

Send a SMS to one or a group of recipients. For details on the input fields, see Intro to SMS. Note: This API is for users with Account Owner scope access token. Users with Extension User scope token should invoke the Extension level Create SMS API with the following definition: POST https://api.phone.com/v4/accounts/:account_id/extensions/:extension_id/sms


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **data** | [**CreateSmsParams**](CreateSmsParams.md)| SMS data | 

### Return type

[**SmsFull**](SmsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountSms**
> SmsFull GetAccountSms($accountId, $smsId)

This service shows the details of an individual SMS.

This service shows the details of an individual SMS. See Intro to SMS for more info on the properties. Note: This API is for users with Account Owner scope access token. Users with Extension User scope token should invoke the Extension level Get SMS API with the following definition: GET https://api.phone.com/v4/accounts/:account_id/extensions/:extension_id/sms/:sms_id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **smsId** | **string**| SMS ID | 

### Return type

[**SmsFull**](SmsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountSms**
> ListSms ListAccountSms($accountId, $filtersId, $filtersFrom, $filtersTo, $filtersDirection, $filtersExtension, $filtersCreatedAt, $sortId, $sortCreatedAt, $limit, $offset, $fields)

Get a list of SMS messages for an account.

Get a list of SMS messages for an account. See Intro to SMS for more info on the properties. Note: This API is for users with Account Owner scope access token. Users with Extension User scope token should invoke the Extension level List SMS API with the following definition: GET https://api.phone.com/v4/accounts/:account_id/extensions/:extension_id/sms


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **filtersId** | [**[]string**](string.md)| ID filter | [optional] 
 **filtersFrom** | **string**| Caller ID filter | [optional] 
 **filtersTo** | **string**| Callee ID filter, the E.164 phone number to send the SMS TO. Note you must encode the + as %2B | [optional] 
 **filtersDirection** | **string**| Direction filter | [optional] 
 **filtersExtension** | [**[]string**](string.md)| Extension filter | [optional] 
 **filtersCreatedAt** | **string**| Date string representing the UTC time that sms was created | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortCreatedAt** | **string**| Sort by created time of message | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListSms**](ListSms.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchAccountSms**
> SmsFull PatchAccountSms($accountId, $smsId, $data)

Update the is_new parameter in a sms record.

Update the is_new parameter in a sms record. See Account SMS for more info on the properties. Note: This API is for users with Account Owner scope access token. Users with Extension User scope token should invoke the Extension level Patch SMS API with the following definition: PATCH https://api.phone.com/v4/accounts/:account_id/extensions/:extension_id/sms/:sms_id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **smsId** | **string**| SMS ID | 
 **data** | [**PatchSmsParams**](PatchSmsParams.md)| Sms data | [optional] 

### Return type

[**SmsFull**](SmsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

