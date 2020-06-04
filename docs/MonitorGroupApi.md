# \MonitorGroupApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MonitorGroupAddMaintenancePeriodToAllMembers**](MonitorGroupApi.md#MonitorGroupAddMaintenancePeriodToAllMembers) | **Post** /MonitorGroup/{monitorGroupGuid}/AddMaintenancePeriodToAllMembers | Adds the provided maintenance period to all monitors in the group specified 
[**MonitorGroupAddMonitorToMonitorGroup**](MonitorGroupApi.md#MonitorGroupAddMonitorToMonitorGroup) | **Post** /MonitorGroup/{monitorGroupGuid}/Members/{monitorGuid} | Adds the specified monitor to the monitor group 
[**MonitorGroupCreateMonitorGroup**](MonitorGroupApi.md#MonitorGroupCreateMonitorGroup) | **Post** /MonitorGroup | Creates a new monitor group
[**MonitorGroupDeleteMonitorGroup**](MonitorGroupApi.md#MonitorGroupDeleteMonitorGroup) | **Delete** /MonitorGroup/{monitorGroupGuid} | Deletes the specified monitor group
[**MonitorGroupGetAllMonitorGroups**](MonitorGroupApi.md#MonitorGroupGetAllMonitorGroups) | **Get** /MonitorGroup | Gets all monitor groups
[**MonitorGroupGetMonitorGroup**](MonitorGroupApi.md#MonitorGroupGetMonitorGroup) | **Get** /MonitorGroup/{monitorGroupGuid} | Gets the details of a monitor group
[**MonitorGroupGetMonitorGroupMembers**](MonitorGroupApi.md#MonitorGroupGetMonitorGroupMembers) | **Get** /MonitorGroup/{monitorGroupGuid}/Members | Gets a list of all members of a monitor group
[**MonitorGroupRemoveMonitorFromMonitorGroup**](MonitorGroupApi.md#MonitorGroupRemoveMonitorFromMonitorGroup) | **Delete** /MonitorGroup/{monitorGroupGuid}/Members/{monitorGuid} | Removes the specified monitor from the monitor group
[**MonitorGroupStartAllMonitorAlertsInGroup**](MonitorGroupApi.md#MonitorGroupStartAllMonitorAlertsInGroup) | **Post** /MonitorGroup/{monitorGroupGuid}/StartAllMonitorAlerts | Starts alerting for all monitors that are a member of the monitor group specified by the monitor group GUID
[**MonitorGroupStartAllMonitorsInGroup**](MonitorGroupApi.md#MonitorGroupStartAllMonitorsInGroup) | **Post** /MonitorGroup/{monitorGroupGuid}/StartAllMonitors | Starts all monitors that are a member of the monitor group specified by the monitor group GUID
[**MonitorGroupStopAllMonitorAlertsInGroup**](MonitorGroupApi.md#MonitorGroupStopAllMonitorAlertsInGroup) | **Post** /MonitorGroup/{monitorGroupGuid}/StopAllMonitorAlerts | Stops alerting for all monitors that are a member of the monitor group specified by the monitor group GUID
[**MonitorGroupStopAllMonitorsInGroup**](MonitorGroupApi.md#MonitorGroupStopAllMonitorsInGroup) | **Post** /MonitorGroup/{monitorGroupGuid}/StopAllMonitors | Stops all monitors that are a member of the monitor group specified by the monitor group GUID
[**MonitorGroupUpdateMonitorGroup**](MonitorGroupApi.md#MonitorGroupUpdateMonitorGroup) | **Put** /MonitorGroup/{monitorGroupGuid} | Updates the monitor group with the Guid specified


# **MonitorGroupAddMaintenancePeriodToAllMembers**
> MonitorGroupAddMaintenancePeriodToAllMembers(ctx, monitorGroupGuid, maintenancePeriod)
Adds the provided maintenance period to all monitors in the group specified 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGroupGuid** | **string**|  | 
  **maintenancePeriod** | [**MaintenancePeriod**](MaintenancePeriod.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGroupAddMonitorToMonitorGroup**
> MonitorGroupAddMonitorToMonitorGroup(ctx, monitorGroupGuid, monitorGuid)
Adds the specified monitor to the monitor group 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGroupGuid** | **string**| The Guid of the monitor group to add the monitor to | 
  **monitorGuid** | **string**| The monitor Guid | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGroupCreateMonitorGroup**
> MonitorGroup MonitorGroupCreateMonitorGroup(ctx, monitorGroup)
Creates a new monitor group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGroup** | [**MonitorGroup**](MonitorGroup.md)| The MonitorGroup object to be created | 

### Return type

[**MonitorGroup**](MonitorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGroupDeleteMonitorGroup**
> MonitorGroupDeleteMonitorGroup(ctx, monitorGroupGuid)
Deletes the specified monitor group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGroupGuid** | **string**| The Guid of the monitor group to be deleted | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGroupGetAllMonitorGroups**
> []MonitorGroup MonitorGroupGetAllMonitorGroups(ctx, )
Gets all monitor groups

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]MonitorGroup**](MonitorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGroupGetMonitorGroup**
> MonitorGroup MonitorGroupGetMonitorGroup(ctx, monitorGroupGuid)
Gets the details of a monitor group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGroupGuid** | **string**| The Guid of the monitor group to be retrieved | 

### Return type

[**MonitorGroup**](MonitorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGroupGetMonitorGroupMembers**
> MonitorGroupMember MonitorGroupGetMonitorGroupMembers(ctx, monitorGroupGuid)
Gets a list of all members of a monitor group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGroupGuid** | **string**| The Guid of the monitor group to retrieve the members for | 

### Return type

[**MonitorGroupMember**](MonitorGroupMember.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGroupRemoveMonitorFromMonitorGroup**
> MonitorGroupRemoveMonitorFromMonitorGroup(ctx, monitorGroupGuid, monitorGuid)
Removes the specified monitor from the monitor group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGroupGuid** | **string**| The Guid of the monitor group to remove the monitor from | 
  **monitorGuid** | **string**| The monitor Guid to be removed | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGroupStartAllMonitorAlertsInGroup**
> MonitorGroupStartAllMonitorAlertsInGroup(ctx, monitorGroupGuid)
Starts alerting for all monitors that are a member of the monitor group specified by the monitor group GUID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGroupGuid** | **string**| The monitor group GUID | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGroupStartAllMonitorsInGroup**
> MonitorGroupStartAllMonitorsInGroup(ctx, monitorGroupGuid)
Starts all monitors that are a member of the monitor group specified by the monitor group GUID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGroupGuid** | **string**| The monitor group GUID | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGroupStopAllMonitorAlertsInGroup**
> MonitorGroupStopAllMonitorAlertsInGroup(ctx, monitorGroupGuid)
Stops alerting for all monitors that are a member of the monitor group specified by the monitor group GUID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGroupGuid** | **string**| The monitor group GUID | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGroupStopAllMonitorsInGroup**
> MonitorGroupStopAllMonitorsInGroup(ctx, monitorGroupGuid)
Stops all monitors that are a member of the monitor group specified by the monitor group GUID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGroupGuid** | **string**| The monitor group GUID | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGroupUpdateMonitorGroup**
> MonitorGroupUpdateMonitorGroup(ctx, item, monitorGroupGuid)
Updates the monitor group with the Guid specified

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **item** | [**MonitorGroup**](MonitorGroup.md)| The monitor group to be updated | 
  **monitorGroupGuid** | **string**| The Guid of the monitor group to be updated | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

