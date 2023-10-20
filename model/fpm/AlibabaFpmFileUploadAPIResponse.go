package fpm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabafpmfileuploadAPIResponse 结算单文件上传 API返回值
// alibaba.fpm.file.upload
//
// 结算单文件上传
type AlibabafpmfileuploadAPIResponse struct {
	model.CommonResponse
	AlibabafpmfileuploadAPIResponseModel
}

// AlibabafpmfileuploadAPIResponseModel is 结算单文件上传 成功返回结果
type AlibabafpmfileuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fpm_file_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabafpmfileuploadResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
