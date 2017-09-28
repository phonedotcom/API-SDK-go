# \QueuesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountQueue**](QueuesApi.md#CreateAccountQueue) | **Post** /accounts/{account_id}/queues | Create a queue.
[**DeleteAccountQueue**](QueuesApi.md#DeleteAccountQueue) | **Delete** /accounts/{account_id}/queues/{queue_id} | Delete a queue.
[**GetAccountQueue**](QueuesApi.md#GetAccountQueue) | **Get** /accounts/{account_id}/queues/{queue_id} | Show details of an individual queue.
[**ListAccountQueues**](QueuesApi.md#ListAccountQueues) | **Get** /accounts/{account_id}/queues | Get a list of queues for an account.
[**ReplaceAccountQueue**](QueuesApi.md#ReplaceAccountQueue) | **Put** /accounts/{account_id}/queues/{queue_id} | Replace a queue.


# **CreateAccountQueue**
> QueueFull CreateAccountQueue($accountId, $data)

Create a queue.

Create a queue. See Account Queues for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **data** | [**CreateQueueParams**](CreateQueueParams.md)| Queue data | [optional] 

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountQueue**
> DeleteEntry DeleteAccountQueue($accountId, $queueId)

Delete a queue.

Delete a queue. See Account Queues for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **queueId** | **int32**| Queue ID | 

### Return type

[**DeleteEntry**](DeleteEntry.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountQueue**
> QueueFull GetAccountQueue($accountId, $queueId)

Show details of an individual queue.

Show details of an individual queue. See Account Queues for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **queueId** | **int32**| Queue ID | 

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountQueues**
> ListQueues ListAccountQueues($accountId, $filtersId, $filtersName, $sortId, $sortName, $limit, $offset, $fields)

Get a list of queues for an account.

Get a list of queues for an account. See Account Queues for more info on the properties.


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

[**ListQueues**](ListQueues.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAccountQueue**
> QueueFull ReplaceAccountQueue($accountId, $queueId, $data)

Replace a queue.

Replace a queue. See Account Queues for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **queueId** | **int32**| Queue ID | 
 **data** | [**CreateQueueParams**](CreateQueueParams.md)| Queue data | [optional] 

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

