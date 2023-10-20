package fpm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFpmFileUploadAPIResponse 结算单文件上传 API返回值
// alibaba.fpm.file.upload
//
// 结算单文件上传
type AlibabaFpmFileUploadAPIResponse struct {
	model.CommonResponse
	AlibabaFpmFileUploadAPIResponseModel
}

// AlibabaFpmFileUploadAPIResponseModel is 结算单文件上传 成功返回结果
type AlibabaFpmFileUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fpm_file_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaFpmFileUploadResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
