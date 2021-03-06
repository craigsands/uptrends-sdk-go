/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

type EscalationLevel struct {
	EscalationMode *EscalationMode `json:"EscalationMode"`
	ThresholdErrorCount int32 `json:"ThresholdErrorCount"`
	ThresholdMinutes int32 `json:"ThresholdMinutes"`
	IsActive bool `json:"IsActive"`
	Message string `json:"Message,omitempty"`
	NumberOfReminders int32 `json:"NumberOfReminders"`
	ReminderDelay int32 `json:"ReminderDelay"`
	IncludeTraceRoute bool `json:"IncludeTraceRoute,omitempty"`
}
