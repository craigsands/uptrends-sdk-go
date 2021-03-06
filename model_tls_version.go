/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends
// TlsVersion : 
type TlsVersion string

// List of TlsVersion
const (
	TLS12_TlsVersion TlsVersion = "Tls12"
	TLS11_TlsVersion TlsVersion = "Tls11"
	TLS10_TlsVersion TlsVersion = "Tls10"
	TLS12_TLS11_TlsVersion TlsVersion = "Tls12_Tls11"
	TLS11_TLS10_TlsVersion TlsVersion = "Tls11_Tls10"
	TLS12_TLS11_TLS10_TlsVersion TlsVersion = "Tls12_Tls11_Tls10"
	TLS12_TLS11_TLS10_WITH_FALLBACK_TlsVersion TlsVersion = "Tls12_Tls11_Tls10_WithFallback"
)
