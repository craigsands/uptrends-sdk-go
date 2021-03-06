# MaintenancePeriod

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique ID of this maintenance period | [default to null]
**ScheduleMode** | [***interface{}**](interface{}.md) | The schedule mode (one time, daily, weekly, monthly) | [default to null]
**StartDateTime** | [**time.Time**](time.Time.md) | The start date/time for this schedule (for one-time schedules only) | [optional] [default to null]
**EndDateTime** | [**time.Time**](time.Time.md) | The end date/time for this maintenance period (for one-time maintenance periods only) | [optional] [default to null]
**WeekDay** | [***interface{}**](interface{}.md) | The weekday for this maintenance period (for weekly maintenance periods only) | [optional] [default to null]
**MonthDay** | **int32** | the month day for this maintenance period (for montly maintenance periods only) | [optional] [default to null]
**StartTime** | **string** | The start time of this maintenance period | [optional] [default to null]
**EndTime** | **string** | The end time of this maintenance period | [optional] [default to null]
**MaintenanceType** | [***interface{}**](interface{}.md) | Indicates whether, during the maintenance periods, only alerting will be disabled, or if the entire monitor will be stopped | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


