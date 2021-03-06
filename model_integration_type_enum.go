/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends
// IntegrationTypeEnum : 
type IntegrationTypeEnum string

// List of IntegrationTypeEnum
const (
	SLACK_IntegrationTypeEnum IntegrationTypeEnum = "Slack"
	PAGER_DUTY_IntegrationTypeEnum IntegrationTypeEnum = "PagerDuty"
	SMS_IntegrationTypeEnum IntegrationTypeEnum = "Sms"
	EMAIL_IntegrationTypeEnum IntegrationTypeEnum = "Email"
	PHONE_IntegrationTypeEnum IntegrationTypeEnum = "Phone"
	STATUSHUB_IntegrationTypeEnum IntegrationTypeEnum = "Statushub"
	GENERIC_WEBHOOK_IntegrationTypeEnum IntegrationTypeEnum = "GenericWebhook"
)
