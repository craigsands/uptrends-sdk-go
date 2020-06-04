# \AlertDefinitionApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertDefinitionAddMonitorGroupToAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionAddMonitorGroupToAlertDefinition) | **Post** /AlertDefinition/{alertDefinitionGuid}/Members/MonitorGroup/{monitorGroupGuid} | Adds a monitor group to the specified alert definition.
[**AlertDefinitionAddMonitorToAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionAddMonitorToAlertDefinition) | **Post** /AlertDefinition/{alertDefinitionGuid}/Members/Monitor/{monitorGuid} | Adds a monitor to the specified alert definition.
[**AlertDefinitionAddOperatorGroupToEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionAddOperatorGroupToEscalationLevel) | **Post** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Members/OperatorGroup/{operatorGroupGuid} | Adds an operator group to the specified escalation level.
[**AlertDefinitionAddOperatorToEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionAddOperatorToEscalationLevel) | **Post** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Members/Operator/{operatorGuid} | Adds an operator to the specified escalation level.
[**AlertDefinitionCreateAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionCreateAlertDefinition) | **Post** /AlertDefinition | Creates a new alert definition.
[**AlertDefinitionDeleteAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionDeleteAlertDefinition) | **Delete** /AlertDefinition/{alertDefinitionGuid} | Deletes an existing alert definition.
[**AlertDefinitionGetAllAlertDefinitions**](AlertDefinitionApi.md#AlertDefinitionGetAllAlertDefinitions) | **Get** /AlertDefinition | Gets a list of all alert definitions.
[**AlertDefinitionGetAllMembers**](AlertDefinitionApi.md#AlertDefinitionGetAllMembers) | **Get** /AlertDefinition/{alertDefinitionGuid}/Members | Gets a list of all monitor and monitor group guids of the specified alert definition.
[**AlertDefinitionGetEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionGetEscalationLevel) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId} | Gets the escalation level information of the specified alert definition.
[**AlertDefinitionGetEscalationLevelIntegration**](AlertDefinitionApi.md#AlertDefinitionGetEscalationLevelIntegration) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration | Gets the integrations for the specified escalation level.
[**AlertDefinitionGetEscalationLevelOperator**](AlertDefinitionApi.md#AlertDefinitionGetEscalationLevelOperator) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Members | Gets the operator and operator group guids for the specified escalation level.
[**AlertDefinitionGetSpecifiedAlertDefinitions**](AlertDefinitionApi.md#AlertDefinitionGetSpecifiedAlertDefinitions) | **Get** /AlertDefinition/{alertDefinitionGuid} | Gets the specified alert definition.
[**AlertDefinitionPatchAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionPatchAlertDefinition) | **Patch** /AlertDefinition/{alertDefinitionGuid} | Partially updates the definition of the specified alert definition.
[**AlertDefinitionPutAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionPutAlertDefinition) | **Put** /AlertDefinition/{alertDefinitionGuid} | Updates the definition of the specified alert definition.
[**AlertDefinitionRemoveMonitorFromAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionRemoveMonitorFromAlertDefinition) | **Delete** /AlertDefinition/{alertDefinitionGuid}/Members/Monitor/{monitorGuid} | Removes a monitor for the specified alert definition.
[**AlertDefinitionRemoveMonitorGroupFromAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionRemoveMonitorGroupFromAlertDefinition) | **Delete** /AlertDefinition/{alertDefinitionGuid}/Members/MonitorGroup/{monitorGroupGuid} | Removes a monitor group for the specified alert definition.
[**AlertDefinitionRemoveOperatorFromEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionRemoveOperatorFromEscalationLevel) | **Delete** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Members/Operator/{operatorGuid} | Removes an operator for the specified escalation level.
[**AlertDefinitionRemoveOperatorGroupFromEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionRemoveOperatorGroupFromEscalationLevel) | **Delete** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Members/OperatorGroup/{operatorGroupGuid} | Removes an operator group for the specified escalation level.
[**AlertDefinitionUpdateIntegrationForEscalationWithPatch**](AlertDefinitionApi.md#AlertDefinitionUpdateIntegrationForEscalationWithPatch) | **Patch** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration/{integrationGuid} | Partially updates an integration to the specified escalation level.
[**AlertDefinitionUpdateIntegrationForEscalationWithPut**](AlertDefinitionApi.md#AlertDefinitionUpdateIntegrationForEscalationWithPut) | **Put** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration/{integrationGuid} | Updates an integration for the specified escalation level.


# **AlertDefinitionAddMonitorGroupToAlertDefinition**
> AlertDefinitionMonitorGroup AlertDefinitionAddMonitorGroupToAlertDefinition(ctx, alertDefinitionGuid, monitorGroupGuid)
Adds a monitor group to the specified alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition to modify. | 
  **monitorGroupGuid** | **string**| The Guid of the monitor group to add. | 

### Return type

[**AlertDefinitionMonitorGroup**](AlertDefinitionMonitorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionAddMonitorToAlertDefinition**
> AlertDefinitionMonitor AlertDefinitionAddMonitorToAlertDefinition(ctx, alertDefinitionGuid, monitorGuid)
Adds a monitor to the specified alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition to modify. | 
  **monitorGuid** | **string**| The Guid of the monitor to add. | 

### Return type

[**AlertDefinitionMonitor**](AlertDefinitionMonitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionAddOperatorGroupToEscalationLevel**
> AlertDefinitionOperatorGroup AlertDefinitionAddOperatorGroupToEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, operatorGroupGuid)
Adds an operator group to the specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 
  **operatorGroupGuid** | **string**| The Guid of the operator group to add. | 

### Return type

[**AlertDefinitionOperatorGroup**](AlertDefinitionOperatorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionAddOperatorToEscalationLevel**
> AlertDefinitionOperator AlertDefinitionAddOperatorToEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, operatorGuid)
Adds an operator to the specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 
  **operatorGuid** | **string**| The Guid of the operator to add. | 

### Return type

[**AlertDefinitionOperator**](AlertDefinitionOperator.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionCreateAlertDefinition**
> AlertDefinition AlertDefinitionCreateAlertDefinition(ctx, alertDefinition)
Creates a new alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinition** | [**AlertDefinition**](AlertDefinition.md)| The details of the alert definition to create. | 

### Return type

[**AlertDefinition**](AlertDefinition.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionDeleteAlertDefinition**
> AlertDefinitionDeleteAlertDefinition(ctx, alertDefinitionGuid)
Deletes an existing alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition to remove. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionGetAllAlertDefinitions**
> []AlertDefinition AlertDefinitionGetAllAlertDefinitions(ctx, )
Gets a list of all alert definitions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AlertDefinition**](AlertDefinition.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionGetAllMembers**
> []AlertDefinitionMember AlertDefinitionGetAllMembers(ctx, alertDefinitionGuid)
Gets a list of all monitor and monitor group guids of the specified alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition for which to return the members. | 

### Return type

[**[]AlertDefinitionMember**](AlertDefinitionMember.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionGetEscalationLevel**
> []EscalationLevel AlertDefinitionGetEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId)
Gets the escalation level information of the specified alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 

### Return type

[**[]EscalationLevel**](EscalationLevel.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionGetEscalationLevelIntegration**
> []Integration AlertDefinitionGetEscalationLevelIntegration(ctx, alertDefinitionGuid, escalationLevelId)
Gets the integrations for the specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 

### Return type

[**[]Integration**](Integration.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionGetEscalationLevelOperator**
> []AlertEscalationLevelMember AlertDefinitionGetEscalationLevelOperator(ctx, alertDefinitionGuid, escalationLevelId)
Gets the operator and operator group guids for the specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 

### Return type

[**[]AlertEscalationLevelMember**](AlertEscalationLevelMember.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionGetSpecifiedAlertDefinitions**
> AlertDefinition AlertDefinitionGetSpecifiedAlertDefinitions(ctx, alertDefinitionGuid)
Gets the specified alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 

### Return type

[**AlertDefinition**](AlertDefinition.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionPatchAlertDefinition**
> AlertDefinitionPatchAlertDefinition(ctx, alertDefinition, alertDefinitionGuid)
Partially updates the definition of the specified alert definition.

This methods accepts parts of an alert definition. Fields that do not require changes can be omitted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinition** | [**AlertDefinition**](AlertDefinition.md)| The partial definition for the alert definition that should be updated. | 
  **alertDefinitionGuid** | **string**| The Guid of the alert definition that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionPutAlertDefinition**
> AlertDefinitionPutAlertDefinition(ctx, alertDefinition, alertDefinitionGuid)
Updates the definition of the specified alert definition.

This methods only accepts a complete alert definition where all fields are specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinition** | [**AlertDefinition**](AlertDefinition.md)| The partial definition for the alert definition that should be updated. | 
  **alertDefinitionGuid** | **string**| The Guid of the alert definition that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionRemoveMonitorFromAlertDefinition**
> AlertDefinitionRemoveMonitorFromAlertDefinition(ctx, alertDefinitionGuid, monitorGuid)
Removes a monitor for the specified alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition to modify. | 
  **monitorGuid** | **string**| The Guid of the monitor to remove. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionRemoveMonitorGroupFromAlertDefinition**
> AlertDefinitionRemoveMonitorGroupFromAlertDefinition(ctx, alertDefinitionGuid, monitorGroupGuid)
Removes a monitor group for the specified alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition to modify. | 
  **monitorGroupGuid** | **string**| The Guid of the monitor group to remove. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionRemoveOperatorFromEscalationLevel**
> AlertDefinitionRemoveOperatorFromEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, operatorGuid)
Removes an operator for the specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 
  **operatorGuid** | **string**| The Guid of the operator to remove. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionRemoveOperatorGroupFromEscalationLevel**
> AlertDefinitionRemoveOperatorGroupFromEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, operatorGroupGuid)
Removes an operator group for the specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 
  **operatorGroupGuid** | **string**| The Guid of the operator group to remove. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionUpdateIntegrationForEscalationWithPatch**
> AlertDefinitionUpdateIntegrationForEscalationWithPatch(ctx, escalationLevelIntegration, alertDefinitionGuid, escalationLevelId, integrationGuid)
Partially updates an integration to the specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **escalationLevelIntegration** | [**EscalationLevelIntegration**](EscalationLevelIntegration.md)| The partial definition for the integration that should be updated. | 
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 
  **integrationGuid** | **string**| The Guid of the integration to update. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionUpdateIntegrationForEscalationWithPut**
> AlertDefinitionUpdateIntegrationForEscalationWithPut(ctx, escalationLevelIntegration, alertDefinitionGuid, escalationLevelId, integrationGuid)
Updates an integration for the specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **escalationLevelIntegration** | [**EscalationLevelIntegration**](EscalationLevelIntegration.md)| The definition for the integration that should be updated. | 
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 
  **integrationGuid** | **string**| The Guid of the integration to update. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

