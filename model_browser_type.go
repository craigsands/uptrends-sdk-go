/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends
// BrowserType : 
type BrowserType string

// List of BrowserType
const (
	IE_BrowserType BrowserType = "IE"
	FIREFOX_BrowserType BrowserType = "Firefox"
	CHROME_BrowserType BrowserType = "Chrome"
	SAFARI_BrowserType BrowserType = "Safari"
	PHANTOM_JS_BrowserType BrowserType = "PhantomJS"
	PHANTOM_JS20_BrowserType BrowserType = "PhantomJS20"
)
