# VaultItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VaultItemGuid** | **string** | The unique key of this vault item | [default to null]
**Hash** | **string** | The hash of this vault item | [optional] [default to null]
**Name** | **string** | The name of this vault item | [optional] [default to null]
**Value** | **string** | The value that is stored in this vault item. Not used for Certificate Archives | [optional] [default to null]
**VaultSectionGuid** | **string** | The unique identifier of the vault section that this vault item belongs to | [default to null]
**VaultItemType** | [***interface{}**](interface{}.md) | The vault item type | [default to null]
**IsSensitive** | **bool** | Whether or not the vault item is considered sensitive.  | [default to null]
**Notes** | **string** | Notes about this vault item | [optional] [default to null]
**UserName** | **string** | The UserName of a credentialset | [optional] [default to null]
**Password** | **string** | The password associated with a credentialset | [optional] [default to null]
**CertificateArchive** | [***interface{}**](interface{}.md) | The certificate archive that is stored in this vault item, if applicable | [optional] [default to null]
**FileInfo** | [***interface{}**](interface{}.md) | The file info that is stored in this vault item, if applicable | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


