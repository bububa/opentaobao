package fpm

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaFpmFileUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaFpmFileUploadAPIResponseModel).Reset()
}

// AlibabaFpmFileUploadAPIResponseModel is 结算单文件上传 成功返回结果
type AlibabaFpmFileUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fpm_file_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaFpmFileUploadResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaFpmFileUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaFpmFileUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaFpmFileUploadAPIResponse)
	},
}

// GetAlibabaFpmFileUploadAPIResponse 从 sync.Pool 获取 AlibabaFpmFileUploadAPIResponse
func GetAlibabaFpmFileUploadAPIResponse() *AlibabaFpmFileUploadAPIResponse {
	return poolAlibabaFpmFileUploadAPIResponse.Get().(*AlibabaFpmFileUploadAPIResponse)
}

// ReleaseAlibabaFpmFileUploadAPIResponse 将 AlibabaFpmFileUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaFpmFileUploadAPIResponse(v *AlibabaFpmFileUploadAPIResponse) {
	v.Reset()
	poolAlibabaFpmFileUploadAPIResponse.Put(v)
}
