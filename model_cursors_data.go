/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

// Cursors can be used to navigate the dataset in a fixed manner
type CursorsData struct {
	// Cursor for next data set
	Next string `json:"Next,omitempty"`
	// Cursor for this data set (data might get updated, depending on your parameters)
	Self string `json:"Self,omitempty"`
}
