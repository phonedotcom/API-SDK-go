# VoicemailFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Voicemail ID. Read-only. | [optional] [default to null]
**DownloadUrl** | **string** | Name. Required. | [optional] [default to null]
**Extension** | [**ExtensionSummary**](ExtensionSummary.md) | Extension where the voicemail is saved into. | [optional] [default to null]
**From** | [**FromObject**](FromObject.md) | The caller&#39;s information | [optional] [default to null]
**To** | **string** | The phone number where the caller is calling | [optional] [default to null]
**IsNew** | **bool** | True when Voicemail is new; False when Voicemail has been listened | [optional] [default to null]
**CreatedAt** | **int32** | Date string representing the UTC time that the object was created in the Phone.com API system. | [optional] [default to null]
**Folder** | **string** | Folder name where voicemail is saved | [optional] [default to null]
**Duration** | **int32** | Length of voicemail in seconds | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


