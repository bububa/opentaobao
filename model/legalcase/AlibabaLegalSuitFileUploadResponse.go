package legalcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
诉讼文件上传接口 APIResponse
alibaba.legal.suit.file.upload

上传文件接口
*/
type AlibabaLegalSuitFileUploadAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_legal_suit_file_upload_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 失败原因
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"