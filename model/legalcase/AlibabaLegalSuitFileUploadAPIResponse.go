package legalcase

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitFileUploadAPIResponse 诉讼文件上传接口 API返回值
// alibaba.legal.suit.file.upload
//
// 上传文件接口
type AlibabaLegalSuitFileUploadAPIResponse struct {
	model.CommonResponse
	AlibabaLegalSuitFileUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalSuitFileUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalSuitFileUploadAPIResponseModel).Reset()
}

// AlibabaLegalSuitFileUploadAPIResponseModel is 诉讼文件上传接口 成功返回结果
type AlibabaLegalSuitFileUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_suit_file_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 失败原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 文件上传后的生成的id
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 失败的code枚举
	CodeError string `json:"code_error,omitempty" xml:"code_error,omitempty"`
	// 是否调用成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalSuitFileUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.Content = ""
	m.CodeError = ""
	m.IsSuccess = false
}

var poolAlibabaLegalSuitFileUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalSuitFileUploadAPIResponse)
	},
}

// GetAlibabaLegalSuitFileUploadAPIResponse 从 sync.Pool 获取 AlibabaLegalSuitFileUploadAPIResponse
func GetAlibabaLegalSuitFileUploadAPIResponse() *AlibabaLegalSuitFileUploadAPIResponse {
	return poolAlibabaLegalSuitFileUploadAPIResponse.Get().(*AlibabaLegalSuitFileUploadAPIResponse)
}

// ReleaseAlibabaLegalSuitFileUploadAPIResponse 将 AlibabaLegalSuitFileUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalSuitFileUploadAPIResponse(v *AlibabaLegalSuitFileUploadAPIResponse) {
	v.Reset()
	poolAlibabaLegalSuitFileUploadAPIResponse.Put(v)
}
