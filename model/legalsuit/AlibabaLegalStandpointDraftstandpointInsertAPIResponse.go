package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointDraftstandpointInsertAPIResponse 编辑后新增草稿口径 API返回值
// alibaba.legal.standpoint.draftstandpoint.insert
//
// 编辑后新增草稿口径
type AlibabaLegalStandpointDraftstandpointInsertAPIResponse struct {
	model.CommonResponse
	AlibabaLegalStandpointDraftstandpointInsertAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointDraftstandpointInsertAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalStandpointDraftstandpointInsertAPIResponseModel).Reset()
}

// AlibabaLegalStandpointDraftstandpointInsertAPIResponseModel is 编辑后新增草稿口径 成功返回结果
type AlibabaLegalStandpointDraftstandpointInsertAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_standpoint_draftstandpoint_insert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 草稿口径id
	Content int64 `json:"content,omitempty" xml:"content,omitempty"`
	// 成功
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointDraftstandpointInsertAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.ErrorCodeRes = 0
	m.Content = 0
	m.SuccessRes = false
}

var poolAlibabaLegalStandpointDraftstandpointInsertAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalStandpointDraftstandpointInsertAPIResponse)
	},
}

// GetAlibabaLegalStandpointDraftstandpointInsertAPIResponse 从 sync.Pool 获取 AlibabaLegalStandpointDraftstandpointInsertAPIResponse
func GetAlibabaLegalStandpointDraftstandpointInsertAPIResponse() *AlibabaLegalStandpointDraftstandpointInsertAPIResponse {
	return poolAlibabaLegalStandpointDraftstandpointInsertAPIResponse.Get().(*AlibabaLegalStandpointDraftstandpointInsertAPIResponse)
}

// ReleaseAlibabaLegalStandpointDraftstandpointInsertAPIResponse 将 AlibabaLegalStandpointDraftstandpointInsertAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalStandpointDraftstandpointInsertAPIResponse(v *AlibabaLegalStandpointDraftstandpointInsertAPIResponse) {
	v.Reset()
	poolAlibabaLegalStandpointDraftstandpointInsertAPIResponse.Put(v)
}
