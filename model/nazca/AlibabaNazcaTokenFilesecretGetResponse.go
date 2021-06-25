package nazca

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取文件秘钥 APIResponse
alibaba.nazca.token.filesecret.get

获取文件秘钥
*/
type AlibabaNazcaTokenFilesecretGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaNazcaTokenFilesecretGetResponse `json:"alibaba_nazca_token_filesecret_get_response,omitempty"`
}

type AlibabaNazcaTokenFilesecretGetResponse struct {

    // error
    Error   string `json:"error,omitempty"`

    // message
    Message   string `json:"message,omitempty"`

    // 文件秘钥
    RetValue   string `json:"ret_value,omitempty"`

    // 错误信息
    SubErrorMessage   string `json:"sub_error_message,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
