package fpm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
结算单文件上传 APIResponse
alibaba.fpm.file.upload

结算单文件上传
*/
type AlibabaFpmFileUploadAPIResponse struct {
    model.CommonResponse
    AlibabaFpmFileUploadResponse
}

type AlibabaFpmFileUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_fpm_file_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaFpmFileUploadResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
