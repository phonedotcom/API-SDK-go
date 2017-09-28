# \ContactsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountExtensionContact**](ContactsApi.md#CreateAccountExtensionContact) | **Post** /accounts/{account_id}/extensions/{extension_id}/contacts | Add a new address book contact for an extension.
[**DeleteAccountExtensionContact**](ContactsApi.md#DeleteAccountExtensionContact) | **Delete** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | Delete a contact from the address book.
[**GetAccountExtensionContact**](ContactsApi.md#GetAccountExtensionContact) | **Get** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | Retrieve the details of an address book contact.
[**ListAccountExtensionContacts**](ContactsApi.md#ListAccountExtensionContacts) | **Get** /accounts/{account_id}/extensions/{extension_id}/contacts | Show the Caller ID options a given extension can use.
[**ReplaceAccountExtensionContact**](ContactsApi.md#ReplaceAccountExtensionContact) | **Put** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | Update the info of a contact in the address book.


# **CreateAccountExtensionContact**
> ContactFull CreateAccountExtensionContact($accountId, $extensionId, $data)

Add a new address book contact for an extension.

Add a new address book contact for an extension. See Account Contacts for more info on the fields in each item.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 
 **data** | [**CreateContactParams**](CreateContactParams.md)| Contact data | [optional] 

### Return type

[**ContactFull**](ContactFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountExtensionContact**
> DeleteEntry DeleteAccountExtensionContact($accountId, $extensionId, $contactId)

Delete a contact from the address book.

Delete a contact from the address book. See Account Contacts for more info on the fields in each item.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 
 **contactId** | **int32**| Contact ID | 

### Return type

[**DeleteEntry**](DeleteEntry.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountExtensionContact**
> ContactFull GetAccountExtensionContact($accountId, $extensionId, $contactId)

Retrieve the details of an address book contact.

Retrieve the details of an address book contact. See Account Contacts for more info on the fields in each item.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 
 **contactId** | **int32**| Contact ID | 

### Return type

[**ContactFull**](ContactFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountExtensionContacts**
> ListContacts ListAccountExtensionContacts($accountId, $extensionId, $filtersId, $filtersGroupId, $filtersUpdatedAt, $sortId, $sortUpdatedAt, $limit, $offset, $fields)

Show the Caller ID options a given extension can use.

Show the Caller ID options a given extension can use. See Intro to Caller IDs for more on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 
 **filtersId** | [**[]string**](string.md)| ID filter | [optional] 
 **filtersGroupId** | [**[]string**](string.md)| Group filter | [optional] 
 **filtersUpdatedAt** | [**[]string**](string.md)| Updated At filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortUpdatedAt** | **string**| Updated At sorting | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListContacts**](ListContacts.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAccountExtensionContact**
> ContactFull ReplaceAccountExtensionContact($accountId, $extensionId, $contactId, $data)

Update the info of a contact in the address book.

Update the info of a contact in the address book. See Account Contacts for more info on the fields in each item.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 
 **contactId** | **int32**| Contact ID | 
 **data** | [**CreateContactParams**](CreateContactParams.md)| Contact data | [optional] 

### Return type

[**ContactFull**](ContactFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

