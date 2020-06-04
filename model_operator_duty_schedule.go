/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

import (
	"time"
)

type OperatorDutySchedule struct {
	// The unique ID of this maintenance period
	Id int32 `json:"Id"`
	// The schedule mode (one time, daily, weekly, monthly)
	ScheduleMode *interface{} `json:"ScheduleMode"`
	// The start date/time for this schedule (for one-time schedules only)
	StartDateTime time.Time `json:"StartDateTime,omitempty"`
	// The end date/time for this maintenance period (for one-time maintenance periods only)
	EndDateTime time.Time `json:"EndDateTime,omitempty"`
	// The weekday for this maintenance period (for weekly maintenance periods only)
	WeekDay *interface{} `json:"WeekDay,omitempty"`
	// the month day for this maintenance period (for montly maintenance periods only)
	MonthDay int32 `json:"MonthDay,omitempty"`
	// The start time of this maintenance period
	StartTime string `json:"StartTime,omitempty"`
	// The end time of this maintenance period
	EndTime string `json:"EndTime,omitempty"`
}
