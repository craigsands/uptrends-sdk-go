# \MonitorCheckApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MonitorCheckGetAccountMonitorChecks**](MonitorCheckApi.md#MonitorCheckGetAccountMonitorChecks) | **Get** /MonitorCheck | Returns all monitor check data.
[**MonitorCheckGetHttpDetails**](MonitorCheckApi.md#MonitorCheckGetHttpDetails) | **Get** /MonitorCheck/{monitorCheckId}/Http | Returns HTTP details for a monitor check.
[**MonitorCheckGetMonitorCheck**](MonitorCheckApi.md#MonitorCheckGetMonitorCheck) | **Get** /MonitorCheck/Monitor/{monitorGuid} | Returns monitor check data for a specific monitor.
[**MonitorCheckGetMonitorGroupData**](MonitorCheckApi.md#MonitorCheckGetMonitorGroupData) | **Get** /MonitorCheck/MonitorGroup/{monitorGroupGuid} | Returns monitor check data for a specific monitor group.
[**MonitorCheckGetMultistepDetails**](MonitorCheckApi.md#MonitorCheckGetMultistepDetails) | **Get** /MonitorCheck/{monitorCheckId}/MultiStepAPI | Returns Multi-Step API details for a monitor check.
[**MonitorCheckGetScreenshots**](MonitorCheckApi.md#MonitorCheckGetScreenshots) | **Get** /MonitorCheck/{monitorCheckId}/Screenshot/{screenshotId} | Gets a specific screenshot for a specified monitor check
[**MonitorCheckGetSingleMonitorCheck**](MonitorCheckApi.md#MonitorCheckGetSingleMonitorCheck) | **Get** /MonitorCheck/{monitorCheckId} | Returns a single monitor check.
[**MonitorCheckGetTransactionDetails**](MonitorCheckApi.md#MonitorCheckGetTransactionDetails) | **Get** /MonitorCheck/{monitorCheckId}/Transaction | Returns transaction step details for a monitor check.
[**MonitorCheckGetWaterfallInfo**](MonitorCheckApi.md#MonitorCheckGetWaterfallInfo) | **Get** /MonitorCheck/{monitorCheckId}/Waterfall | Returns waterfall information for a monitor check.


# **MonitorCheckGetAccountMonitorChecks**
> MonitorCheckResponse MonitorCheckGetAccountMonitorChecks(ctx, optional)
Returns all monitor check data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MonitorCheckGetAccountMonitorChecksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorCheckGetAccountMonitorChecksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **errorLevel** | **optional.String**| Error level filter that should be applied. (default &#x3D; NoError and above) | 
 **cursor** | **optional.String**| A cursor value that should be used for traversing the dataset. | 
 **sorting** | **optional.String**| Sorting direction based on timestamp. | [default to Descending]
 **take** | **optional.Int32**| The number of records to return (Max value &#x3D; 100) | [default to 100]
 **start** | **optional.Time**| The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | **optional.Time**| The end of a custom period | 
 **presetPeriod** | **optional.String**| The requested time period. | [default to Last24Hours]

### Return type

[**MonitorCheckResponse**](MonitorCheckResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorCheckGetHttpDetails**
> HttpDetailsResponse MonitorCheckGetHttpDetails(ctx, monitorCheckId)
Returns HTTP details for a monitor check.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorCheckId** | **int64**| The monitor check Id to get the detailed data for. | 

### Return type

[**HttpDetailsResponse**](HttpDetailsResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorCheckGetMonitorCheck**
> MonitorCheckResponse MonitorCheckGetMonitorCheck(ctx, monitorGuid, optional)
Returns monitor check data for a specific monitor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The Guid of the monitor to get monitor checks for. | 
 **optional** | ***MonitorCheckGetMonitorCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorCheckGetMonitorCheckOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **errorLevel** | **optional.String**| Error level filter that should be applied. (default &#x3D; NoError and above) | 
 **cursor** | **optional.String**| A cursor value that should be used for traversing the dataset. | 
 **sorting** | **optional.String**| Sorting direction based on timestamp. | [default to Descending]
 **take** | **optional.Int32**| The number of records to return (Max value &#x3D; 100) | [default to 100]
 **start** | **optional.Time**| The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | **optional.Time**| The end of a custom period | 
 **presetPeriod** | **optional.String**| The requested time period. | [default to Last24Hours]

### Return type

[**MonitorCheckResponse**](MonitorCheckResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorCheckGetMonitorGroupData**
> MonitorCheckResponse MonitorCheckGetMonitorGroupData(ctx, monitorGroupGuid, optional)
Returns monitor check data for a specific monitor group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGroupGuid** | **string**| The Guid of the monitor group to get monitor checks for. | 
 **optional** | ***MonitorCheckGetMonitorGroupDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorCheckGetMonitorGroupDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **errorLevel** | **optional.String**| Error level filter that should be applied. (default &#x3D; NoError and above) | 
 **cursor** | **optional.String**| A cursor value that should be used for traversing the dataset. | 
 **sorting** | **optional.String**| Sorting direction based on timestamp. | [default to Descending]
 **take** | **optional.Int32**| The number of records to return (Max value &#x3D; 100) | [default to 100]
 **start** | **optional.Time**| The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | **optional.Time**| The end of a custom period | 
 **presetPeriod** | **optional.String**| The requested time period. | [default to Last24Hours]

### Return type

[**MonitorCheckResponse**](MonitorCheckResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorCheckGetMultistepDetails**
> MsaDetailsResponse MonitorCheckGetMultistepDetails(ctx, monitorCheckId)
Returns Multi-Step API details for a monitor check.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorCheckId** | **int64**| The monitor check Id to get the detailed data for. | 

### Return type

[**MsaDetailsResponse**](MsaDetailsResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorCheckGetScreenshots**
> ScreenshotResponse MonitorCheckGetScreenshots(ctx, monitorCheckId, screenshotId)
Gets a specific screenshot for a specified monitor check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorCheckId** | **int64**| The monitor check Id to get the screenshot data for. | 
  **screenshotId** | **string**| The screenshot Id of the screenshot to get. | 

### Return type

[**ScreenshotResponse**](ScreenshotResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorCheckGetSingleMonitorCheck**
> SingleMonitorCheckResponse MonitorCheckGetSingleMonitorCheck(ctx, monitorCheckId)
Returns a single monitor check.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorCheckId** | **int64**| The Id of the monitor check to get the data for. | 

### Return type

[**SingleMonitorCheckResponse**](SingleMonitorCheckResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorCheckGetTransactionDetails**
> TransactionDetailsResponse MonitorCheckGetTransactionDetails(ctx, monitorCheckId)
Returns transaction step details for a monitor check.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorCheckId** | **int64**| The monitor check Id to get the detailed data for. | 

### Return type

[**TransactionDetailsResponse**](TransactionDetailsResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorCheckGetWaterfallInfo**
> WaterfallResponse MonitorCheckGetWaterfallInfo(ctx, monitorCheckId, optional)
Returns waterfall information for a monitor check.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorCheckId** | **int64**| The monitor check Id to get the detailed data for. | 
 **optional** | ***MonitorCheckGetWaterfallInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorCheckGetWaterfallInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **step** | **optional.Int32**| For transaction waterfalls only: the transaction step to get the waterfall for. | 

### Return type

[**WaterfallResponse**](WaterfallResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

