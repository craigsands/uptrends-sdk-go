# \MonitorApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MonitorCleanupMaintenancePeriods**](MonitorApi.md#MonitorCleanupMaintenancePeriods) | **Post** /Monitor/{monitorGuid}/MaintenancePeriod/Cleanup/{beforeDate} | Clears out all one-time maintenance periods for the specified monitor older than the specified date
[**MonitorCloneMonitor**](MonitorApi.md#MonitorCloneMonitor) | **Post** /Monitor/{monitorGuid}/Clone | Creates a clone (duplicate) of the specified monitor.
[**MonitorCreateMaintenancePeriodForMonitor**](MonitorApi.md#MonitorCreateMaintenancePeriodForMonitor) | **Post** /Monitor/{monitorGuid}/MaintenancePeriod | Saves the new maintenance period provided for the specified monitor
[**MonitorDeleteMaintenancePeriodFromMonitor**](MonitorApi.md#MonitorDeleteMaintenancePeriodFromMonitor) | **Delete** /Monitor/{monitorGuid}/MaintenancePeriod/{maintenancePeriodId} | Deletes the specified maintenance period from the specified monitor
[**MonitorDeleteMonitor**](MonitorApi.md#MonitorDeleteMonitor) | **Delete** /Monitor/{monitorGuid} | Deletes the specified monitor.
[**MonitorGetAllMaintenancePeriodsForMonitor**](MonitorApi.md#MonitorGetAllMaintenancePeriodsForMonitor) | **Get** /Monitor/{monitorGuid}/MaintenancePeriod | Finds all maintenance periods for a monitor.
[**MonitorGetMonitor**](MonitorApi.md#MonitorGetMonitor) | **Get** /Monitor/{monitorGuid} | Returns the definition of the specified monitor. 
[**MonitorGetMonitors**](MonitorApi.md#MonitorGetMonitors) | **Get** /Monitor | Returns the definition of all monitors available in the account.
[**MonitorPatchMonitor**](MonitorApi.md#MonitorPatchMonitor) | **Patch** /Monitor/{monitorGuid} | Partially updates the definition of the specified monitor.
[**MonitorPostMonitor**](MonitorApi.md#MonitorPostMonitor) | **Post** /Monitor | Creates a new monitor.
[**MonitorPutMonitor**](MonitorApi.md#MonitorPutMonitor) | **Put** /Monitor/{monitorGuid} | Updates the definition of the specified monitor.
[**MonitorUpdateMaintenancePeriodForMonitor**](MonitorApi.md#MonitorUpdateMaintenancePeriodForMonitor) | **Put** /Monitor/{monitorGuid}/MaintenancePeriod/{maintenancePeriodId} | Updates the specified maintenance schedule for the specified monitor


# **MonitorCleanupMaintenancePeriods**
> MonitorCleanupMaintenancePeriods(ctx, monitorGuid, beforeDate)
Clears out all one-time maintenance periods for the specified monitor older than the specified date

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**|  | 
  **beforeDate** | **time.Time**| A string representing the date, formatted as \&quot;yyyy-MM-dd\&quot; | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorCloneMonitor**
> Monitor MonitorCloneMonitor(ctx, monitorGuid, optional)
Creates a clone (duplicate) of the specified monitor.

Upon creation, the new monitor will be inactive. This allows you to make the necessary changes before you activate it. All other settings will be transferred to the new monitor as-is.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The guid of the monitor you want to clone. | 
 **optional** | ***MonitorCloneMonitorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorCloneMonitorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeMaintenancePeriods** | **optional.Bool**| Whether or not to also copy the maintenance periods into the clone. | [default to true]
 **includeMonitorGroups** | **optional.Bool**| Whether or not to also copy the monitor group memberships into the clone. | [default to true]

### Return type

[**Monitor**](Monitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorCreateMaintenancePeriodForMonitor**
> MaintenancePeriod MonitorCreateMaintenancePeriodForMonitor(ctx, maintenancePeriod, monitorGuid)
Saves the new maintenance period provided for the specified monitor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **maintenancePeriod** | [**MaintenancePeriod**](MaintenancePeriod.md)|  | 
  **monitorGuid** | **string**|  | 

### Return type

[**MaintenancePeriod**](MaintenancePeriod.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorDeleteMaintenancePeriodFromMonitor**
> MonitorDeleteMaintenancePeriodFromMonitor(ctx, monitorGuid, maintenancePeriodId)
Deletes the specified maintenance period from the specified monitor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**|  | 
  **maintenancePeriodId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorDeleteMonitor**
> MonitorDeleteMonitor(ctx, monitorGuid)
Deletes the specified monitor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The guid of the monitor you want to delete. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGetAllMaintenancePeriodsForMonitor**
> []MaintenancePeriod MonitorGetAllMaintenancePeriodsForMonitor(ctx, monitorGuid)
Finds all maintenance periods for a monitor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The guid of the monitor you want to find the maintenance periods of. | 

### Return type

[**[]MaintenancePeriod**](MaintenancePeriod.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGetMonitor**
> Monitor MonitorGetMonitor(ctx, monitorGuid, optional)
Returns the definition of the specified monitor. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The Guid of the requested monitor. | 
 **optional** | ***MonitorGetMonitorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorGetMonitorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Provide the option to only retrieve the requested fields. E.g. \&quot;Name,IsActive\&quot;. | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGetMonitors**
> []Monitor MonitorGetMonitors(ctx, optional)
Returns the definition of all monitors available in the account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MonitorGetMonitorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorGetMonitorsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Provide the option to only retrieve the requested fields. E.g. \&quot;Name,IsActive\&quot;. | 

### Return type

[**[]Monitor**](Monitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorPatchMonitor**
> MonitorPatchMonitor(ctx, monitor, monitorGuid)
Partially updates the definition of the specified monitor.

This methods accepts parts of a monitor definition. We recommend retrieving the existing definition first (using the GET method). You can then process the changes you want to make and send back these changes only using this PATCH method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitor** | [**Monitor**](Monitor.md)| The partial definition for the monitor that should be updated. | 
  **monitorGuid** | **string**| The Guid of the monitor that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorPostMonitor**
> Monitor MonitorPostMonitor(ctx, monitor)
Creates a new monitor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitor** | [**Monitor**](Monitor.md)| The complete definition of the monitor that should be created. | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorPutMonitor**
> MonitorPutMonitor(ctx, monitor, monitorGuid)
Updates the definition of the specified monitor.

This methods only accepts a complete monitor definition. We recommend retrieving the existing definition first (using the GET method). You can then process the changes you want to make and send back the updated definition using this PUT method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitor** | [**Monitor**](Monitor.md)| The complete definition for the monitor that should be updated. | 
  **monitorGuid** | **string**| The Guid of the monitor that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorUpdateMaintenancePeriodForMonitor**
> MonitorUpdateMaintenancePeriodForMonitor(ctx, monitorGuid, maintenancePeriodId, maintenancePeriod)
Updates the specified maintenance schedule for the specified monitor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**|  | 
  **maintenancePeriodId** | **int32**|  | 
  **maintenancePeriod** | [**MaintenancePeriod**](MaintenancePeriod.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

