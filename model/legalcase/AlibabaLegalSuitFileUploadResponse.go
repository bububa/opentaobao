package legalcase

import (
    "github.com/bububa/opentaobao/model"
)

/* 
诉讼文件上传接口 APIResponse
alibaba.legal.suit.file.upload

上传文件接口
*/
type AlibabaLegalSuitFileUploadAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaLegalSuitFileUploadResponse `json:"alibaba_legal_suit_file_upload_response,omitempty"` 
    AlibabaLegalSuitFileUploadResponse
}

/* model for simplify = false
type AlibabaLegalSuitFileUploadResponse struct {

    // 失败原因
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 是否调用成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 文件上传后的生成的id
    
    Content   string `json:"content,omitempty"`
    

    // 失败的code枚举
    
    CodeError   string `json:"code_error,omitempty"`
    

}
*/

type AlibabaLegalSuitFileUploadResponse struct {

    // 失败原因
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 是否调用成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 文件上传后的生成的id
    Content   string `json:"content,omitempty"`

    // 失败的code枚举
    CodeError   string `json:"code_error,omitempty"`

}
