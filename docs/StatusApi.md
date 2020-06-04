# \StatusApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatusGetMonitorGroupStatus**](StatusApi.md#StatusGetMonitorGroupStatus) | **Get** /Status/MonitorGroup/{monitorGroupGuid} | Gets a list of all monitor group status data.
[**StatusGetMonitorStatus**](StatusApi.md#StatusGetMonitorStatus) | **Get** /Status/Monitor/{monitorGuid} | Gets all monitor status data.


# **StatusGetMonitorGroupStatus**
> MonitorStatusListResponse StatusGetMonitorGroupStatus(ctx, monitorGroupGuid)
Gets a list of all monitor group status data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGroupGuid** | **string**| The Guid of the monitor group. | 

### Return type

[**MonitorStatusListResponse**](MonitorStatusListResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatusGetMonitorStatus**
> MonitorStatusResponse StatusGetMonitorStatus(ctx, monitorGuid)
Gets all monitor status data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The Guid of the monitor. | 

### Return type

[**MonitorStatusResponse**](MonitorStatusResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

