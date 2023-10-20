package legalcase

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseCourtTimeUpdateAPIResponse 开庭时间变更 API返回值
// alibaba.legal.case.court.time.update
//
// 修改案件的开庭时间
type AlibabaLegalCaseCourtTimeUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaLegalCaseCourtTimeUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalCaseCourtTimeUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalCaseCourtTimeUpdateAPIResponseModel).Reset()
}

// AlibabaLegalCaseCourtTimeUpdateAPIResponseModel is 开庭时间变更 成功返回结果
type AlibabaLegalCaseCourtTimeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_case_court_time_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalCaseCourtTimeUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLegalCaseCourtTimeUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalCaseCourtTimeUpdateAPIResponse)
	},
}

// GetAlibabaLegalCaseCourtTimeUpdateAPIResponse 从 sync.Pool 获取 AlibabaLegalCaseCourtTimeUpdateAPIResponse
func GetAlibabaLegalCaseCourtTimeUpdateAPIResponse() *AlibabaLegalCaseCourtTimeUpdateAPIResponse {
	return poolAlibabaLegalCaseCourtTimeUpdateAPIResponse.Get().(*AlibabaLegalCaseCourtTimeUpdateAPIResponse)
}

// ReleaseAlibabaLegalCaseCourtTimeUpdateAPIResponse 将 AlibabaLegalCaseCourtTimeUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalCaseCourtTimeUpdateAPIResponse(v *AlibabaLegalCaseCourtTimeUpdateAPIResponse) {
	v.Reset()
	poolAlibabaLegalCaseCourtTimeUpdateAPIResponse.Put(v)
}
