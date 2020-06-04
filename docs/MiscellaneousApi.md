# \MiscellaneousApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MiscellaneousGetAllOutgoingPhoneNumbers**](MiscellaneousApi.md#MiscellaneousGetAllOutgoingPhoneNumbers) | **Get** /OutgoingPhoneNumber | Gets a list of all outgoing phone numbers available.
[**MiscellaneousGetAllTimezones**](MiscellaneousApi.md#MiscellaneousGetAllTimezones) | **Get** /Timezone | Gets all timezones available.
[**MiscellaneousGetTimezoneById**](MiscellaneousApi.md#MiscellaneousGetTimezoneById) | **Get** /Timezone/{timezoneId} | Gets the timezone with the specified Id.


# **MiscellaneousGetAllOutgoingPhoneNumbers**
> []OutgoingPhoneNumberDetails MiscellaneousGetAllOutgoingPhoneNumbers(ctx, )
Gets a list of all outgoing phone numbers available.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]OutgoingPhoneNumberDetails**](OutgoingPhoneNumberDetails.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MiscellaneousGetAllTimezones**
> []Timezone MiscellaneousGetAllTimezones(ctx, )
Gets all timezones available.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Timezone**](Timezone.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MiscellaneousGetTimezoneById**
> Timezone MiscellaneousGetTimezoneById(ctx, timezoneId)
Gets the timezone with the specified Id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timezoneId** | **int32**|  | 

### Return type

[**Timezone**](Timezone.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

