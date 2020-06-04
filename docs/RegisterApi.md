# \RegisterApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegisterPost**](RegisterApi.md#RegisterPost) | **Post** /Register | Creates a new API account.


# **RegisterPost**
> RegistrationResponse RegisterPost(ctx, )
Creates a new API account.

This method requires that you specify the username and password of an Uptrends operator login as authentication. When registration is successful, please save the UserName and Password fields from the response; these are your new API credentials.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RegistrationResponse**](RegistrationResponse.md)

### Authorization

[user-basicauth](../README.md#user-basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

