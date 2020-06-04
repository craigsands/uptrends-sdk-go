/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

type VaultItem struct {
	// The unique key of this vault item
	VaultItemGuid string `json:"VaultItemGuid"`
	// The hash of this vault item
	Hash string `json:"Hash,omitempty"`
	// The name of this vault item
	Name string `json:"Name,omitempty"`
	// The value that is stored in this vault item. Not used for Certificate Archives
	Value string `json:"Value,omitempty"`
	// The unique identifier of the vault section that this vault item belongs to
	VaultSectionGuid string `json:"VaultSectionGuid"`
	// The vault item type
	VaultItemType *interface{} `json:"VaultItemType"`
	// Whether or not the vault item is considered sensitive. 
	IsSensitive bool `json:"IsSensitive"`
	// Notes about this vault item
	Notes string `json:"Notes,omitempty"`
	// The UserName of a credentialset
	UserName string `json:"UserName,omitempty"`
	// The password associated with a credentialset
	Password string `json:"Password,omitempty"`
	// The certificate archive that is stored in this vault item, if applicable
	CertificateArchive *interface{} `json:"CertificateArchive,omitempty"`
	// The file info that is stored in this vault item, if applicable
	FileInfo *interface{} `json:"FileInfo,omitempty"`
}