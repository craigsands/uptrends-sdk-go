# MonitorCheckAttributes

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorGuid** | **string** | Monitor identifier | [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | Date/time stamp of the check | [default to null]
**ErrorCode** | **int32** | The numeric Uptrends error code in case of an error result, or 0 in case of an OK result. | [default to null]
**TotalTime** | **float64** | The number of milliseconds needed to complete the monitor check. | [default to null]
**ResolveTime** | **float64** | The number of milliseconds needed to perform the DNS query for this check, when appropriate. | [default to null]
**ConnectionTime** | **float64** | The number of milliseconds needed to establish a connection. | [default to null]
**DownloadTime** | **float64** | The number of milliseconds needed to download the response data. | [default to null]
**TotalBytes** | **int32** | The number of downloaded bytes for this check. | [optional] [default to null]
**ResolvedIpAddress** | **string** | The IP address that was found for the specified domain name as part of this monitor check. | [optional] [default to null]
**ErrorLevel** | [***interface{}**](interface{}.md) | A value that represents the OK/Error state for this check: NoError if the result was OK, Unconfirmed if an error was found, Confirmed if an error was found as a double check, right after an Unconfirmed error. | [default to null]
**ErrorDescription** | **string** | A text value that describes the error that was found, or OK if no error was found. | [optional] [default to null]
**ErrorMessage** | **string** | Any additional error information, if available. | [optional] [default to null]
**StagingMode** | **bool** | Did the check come from a staging monitor? | [default to null]
**ServerId** | **int32** | The Id of the Uptrends checkpoint server that performed this check. | [default to null]
**HttpStatusCode** | **int32** | The HTTP status code returned (if applicable) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


