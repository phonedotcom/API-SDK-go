# \DevicesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountDevice**](DevicesApi.md#CreateAccountDevice) | **Post** /accounts/{account_id}/devices | Register a generic VoIP device.
[**DeleteAccountDevice**](DevicesApi.md#DeleteAccountDevice) | **Delete** /accounts/{account_id}/devices/{device_id} | Delete a VoIP device.
[**GetAccountDevice**](DevicesApi.md#GetAccountDevice) | **Get** /accounts/{account_id}/devices/{device_id} | Show details of an individual VoIP device.
[**ListAccountDevices**](DevicesApi.md#ListAccountDevices) | **Get** /accounts/{account_id}/devices | Get a list of VoIP devices associated with your account.
[**ReplaceAccountDevice**](DevicesApi.md#ReplaceAccountDevice) | **Put** /accounts/{account_id}/devices/{device_id} | Update the details of an individual VoIP device.


# **CreateAccountDevice**
> DeviceFull CreateAccountDevice($accountId, $data)

Register a generic VoIP device.

Register a generic VoIP device. See Devices for more detail.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **data** | [**CreateDeviceParams**](CreateDeviceParams.md)| Device data | [optional] 

### Return type

[**DeviceFull**](DeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountDevice**
> DeleteEntry DeleteAccountDevice($accountId, $deviceId)

Delete a VoIP device.

Delete a VoIP device. See Devices for more detail.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **deviceId** | **int32**| Device ID | 

### Return type

[**DeleteEntry**](DeleteEntry.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountDevice**
> DeviceFull GetAccountDevice($accountId, $deviceId)

Show details of an individual VoIP device.

Show details of an individual VoIP device. See Devices for more detail.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **deviceId** | **int32**| Device ID | 

### Return type

[**DeviceFull**](DeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountDevices**
> ListDevices ListAccountDevices($accountId, $filtersId, $filtersName, $sortId, $sortName, $limit, $offset, $fields)

Get a list of VoIP devices associated with your account.

Get a list of VoIP devices associated with your account. See Devices for more detail.


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

[**ListDevices**](ListDevices.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAccountDevice**
> DeviceFull ReplaceAccountDevice($accountId, $deviceId, $data)

Update the details of an individual VoIP device.

Update the details of an individual VoIP device. See Devices for more detail.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **deviceId** | **int32**| Device ID | 
 **data** | [**CreateDeviceParams**](CreateDeviceParams.md)| Device data | [optional] 

### Return type

[**DeviceFull**](DeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

