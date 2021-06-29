package nazca

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取文件秘钥 API返回值 
alibaba.nazca.token.filesecret.get

获取文件秘钥
*/
type AlibabaNazcaTokenFilesecretGetAPIResponse struct {
    model.CommonResponse
    AlibabaNazcaTokenFilesecretGetResponse
}

// 获取文件秘钥 成功返回结果
type AlibabaNazcaTokenFilesecretGetResponse struct {
    XMLName xml.Name `xml:"alibaba_nazca_token_filesecret_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // error
    Error   string `json:"error,omitempty" xml:"error,omitempty"`
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 文件秘钥
    RetValue   string `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
    // 错误信息
    SubErrorMessage   string `json:"sub_error_message,omitempty" xml:"sub_error_message,omitempty"`
    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
