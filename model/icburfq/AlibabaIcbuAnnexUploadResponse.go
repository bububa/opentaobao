package icburfq

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上传附件获取附件files_str API返回值 
alibaba.icbu.annex.upload

上传附件获取附件files_str
*/
type AlibabaIcbuAnnexUploadAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuAnnexUploadResponse
}

// 上传附件获取附件files_str 成功返回结果
type AlibabaIcbuAnnexUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_annex_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回错误码
    ErrType   string `json:"err_type,omitempty" xml:"err_type,omitempty"`
    // 错误信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 文件file_str
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
