package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointInsertdraftAPIResponse 插入草稿 API返回值
// alibaba.legal.standpoint.insertdraft
//
// 插入草稿
type AlibabaLegalStandpointInsertdraftAPIResponse struct {
	model.CommonResponse
	AlibabaLegalStandpointInsertdraftAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointInsertdraftAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalStandpointInsertdraftAPIResponseModel).Reset()
}

// AlibabaLegalStandpointInsertdraftAPIResponseModel is 插入草稿 成功返回结果
type AlibabaLegalStandpointInsertdraftAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_standpoint_insertdraft_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 状态code
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 返回内容
	Content int64 `json:"content,omitempty" xml:"content,omitempty"`
	// 是否调用成功
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointInsertdraftAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.ErrorCodeRes = 0
	m.Content = 0
	m.SuccessRes = false
}

var poolAlibabaLegalStandpointInsertdraftAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalStandpointInsertdraftAPIResponse)
	},
}

// GetAlibabaLegalStandpointInsertdraftAPIResponse 从 sync.Pool 获取 AlibabaLegalStandpointInsertdraftAPIResponse
func GetAlibabaLegalStandpointInsertdraftAPIResponse() *AlibabaLegalStandpointInsertdraftAPIResponse {
	return poolAlibabaLegalStandpointInsertdraftAPIResponse.Get().(*AlibabaLegalStandpointInsertdraftAPIResponse)
}

// ReleaseAlibabaLegalStandpointInsertdraftAPIResponse 将 AlibabaLegalStandpointInsertdraftAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalStandpointInsertdraftAPIResponse(v *AlibabaLegalStandpointInsertdraftAPIResponse) {
	v.Reset()
	poolAlibabaLegalStandpointInsertdraftAPIResponse.Put(v)
}
