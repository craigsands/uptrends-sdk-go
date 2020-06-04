/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends
// ScheduleMode : 
type ScheduleMode string

// List of ScheduleMode
const (
	ONE_TIME_ScheduleMode ScheduleMode = "OneTime"
	DAILY_ScheduleMode ScheduleMode = "Daily"
	WEEKLY_ScheduleMode ScheduleMode = "Weekly"
	MONTHLY_ScheduleMode ScheduleMode = "Monthly"
)