/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

type RegistrationResponse struct {
	// The user name of the new API account
	UserName string `json:"UserName,omitempty"`
	// The password of the new API account
	Password string `json:"Password,omitempty"`
	// The Uptrends Account Id for which the new API account was created
	AccountId string `json:"AccountId,omitempty"`
	// The Uptrends Operator on behalf of whom the new API account was created
	OperatorName string `json:"OperatorName,omitempty"`
	Status *RegisterStatus `json:"status"`
	Message string `json:"message,omitempty"`
}