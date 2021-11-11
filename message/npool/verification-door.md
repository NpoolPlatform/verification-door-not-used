# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/verification-door.proto](#npool/verification-door.proto)
    - [DeleteUserGoogleAuthRequest](#verification.door.v1.DeleteUserGoogleAuthRequest)
    - [DeleteUserGoogleAuthResponse](#verification.door.v1.DeleteUserGoogleAuthResponse)
    - [GetQRcodeURLRequest](#verification.door.v1.GetQRcodeURLRequest)
    - [GetQRcodeURLResponse](#verification.door.v1.GetQRcodeURLResponse)
    - [SendEmailRequest](#verification.door.v1.SendEmailRequest)
    - [SendEmailResponse](#verification.door.v1.SendEmailResponse)
    - [SendSmsRequest](#verification.door.v1.SendSmsRequest)
    - [SendSmsResponse](#verification.door.v1.SendSmsResponse)
    - [VerifyCodeRequest](#verification.door.v1.VerifyCodeRequest)
    - [VerifyCodeResponse](#verification.door.v1.VerifyCodeResponse)
    - [VerifyGoogleAuthRequest](#verification.door.v1.VerifyGoogleAuthRequest)
    - [VerifyGoogleAuthResponse](#verification.door.v1.VerifyGoogleAuthResponse)
    - [VersionResponse](#verification.door.v1.VersionResponse)
  
    - [VerificationDoor](#verification.door.v1.VerificationDoor)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/verification-door.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/verification-door.proto



<a name="verification.door.v1.DeleteUserGoogleAuthRequest"></a>

### DeleteUserGoogleAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |






<a name="verification.door.v1.DeleteUserGoogleAuthResponse"></a>

### DeleteUserGoogleAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="verification.door.v1.GetQRcodeURLRequest"></a>

### GetQRcodeURLRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Username | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="verification.door.v1.GetQRcodeURLResponse"></a>

### GetQRcodeURLResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="verification.door.v1.SendEmailRequest"></a>

### SendEmailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| Email | [string](#string) |  |  |
| Intention | [string](#string) |  |  |






<a name="verification.door.v1.SendEmailResponse"></a>

### SendEmailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="verification.door.v1.SendSmsRequest"></a>

### SendSmsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| Phone | [string](#string) |  |  |
| Intention | [string](#string) |  |  |






<a name="verification.door.v1.SendSmsResponse"></a>

### SendSmsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="verification.door.v1.VerifyCodeRequest"></a>

### VerifyCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| Code | [string](#string) |  |  |
| VerifyType | [string](#string) |  |  |






<a name="verification.door.v1.VerifyCodeResponse"></a>

### VerifyCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="verification.door.v1.VerifyGoogleAuthRequest"></a>

### VerifyGoogleAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| Code | [string](#string) |  |  |






<a name="verification.door.v1.VerifyGoogleAuthResponse"></a>

### VerifyGoogleAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="verification.door.v1.VersionResponse"></a>

### VersionResponse
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="verification.door.v1.VerificationDoor"></a>

### VerificationDoor
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#verification.door.v1.VersionResponse) | Method Version |
| GetQRcodeURL | [GetQRcodeURLRequest](#verification.door.v1.GetQRcodeURLRequest) | [GetQRcodeURLResponse](#verification.door.v1.GetQRcodeURLResponse) |  |
| VerifyGoogleAuth | [VerifyGoogleAuthRequest](#verification.door.v1.VerifyGoogleAuthRequest) | [VerifyGoogleAuthResponse](#verification.door.v1.VerifyGoogleAuthResponse) |  |
| DeleteUserGoogleAuth | [DeleteUserGoogleAuthRequest](#verification.door.v1.DeleteUserGoogleAuthRequest) | [DeleteUserGoogleAuthResponse](#verification.door.v1.DeleteUserGoogleAuthResponse) |  |
| SendEmail | [SendEmailRequest](#verification.door.v1.SendEmailRequest) | [SendEmailResponse](#verification.door.v1.SendEmailResponse) |  |
| SendSms | [SendSmsRequest](#verification.door.v1.SendSmsRequest) | [SendSmsResponse](#verification.door.v1.SendSmsResponse) |  |
| VerifyCode | [VerifyCodeRequest](#verification.door.v1.VerifyCodeRequest) | [VerifyCodeResponse](#verification.door.v1.VerifyCodeResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
