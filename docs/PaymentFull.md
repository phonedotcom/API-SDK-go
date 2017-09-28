# PaymentFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Payment Method ID | [optional] [default to null]
**VoipId** | **int32** | API Account ID | [optional] [default to null]
**Status** | **string** | &#39;primary&#39; &#x3D; primary card used for billing; &#39;onfile&#39; &#x3D; card on file; &#39;hidden&#39; &#x3D; deleted card; | [optional] [default to null]
**Nickname** | **string** | Name of Card | [optional] [default to null]
**Type_** | **string** | &#39;cc&#39; for credit card | [optional] [default to null]
**CreatedAt** | **int32** | Time payment record is created | [optional] [default to null]
**Contact** | [**ContactResponse**](ContactResponse.md) | Detail of contact person | [optional] [default to null]
**DeclineCount** | **int32** | Number of times the payment method was declined | [optional] [default to null]
**NextChargeDate** | **int32** | Next billing date | [optional] [default to null]
**UpdatedAt** | **int32** | Last time the payment method was updated | [optional] [default to null]
**CcToken** | **string** | Encrypted credit card token to be used for billing | [optional] [default to null]
**CcNumber** | **string** | Credit card number partially masked with xxxxxxxx | [optional] [default to null]
**CcExp** | **string** | Credit card expiration date | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


