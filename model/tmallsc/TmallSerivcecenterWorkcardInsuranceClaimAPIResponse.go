package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallSerivcecenterWorkcardInsuranceClaimAPIResponse 保险理赔回传工单记录 API返回值
// tmall.serivcecenter.workcard.insurance.claim
//
// 保险理赔回传工单记录
type TmallSerivcecenterWorkcardInsuranceClaimAPIResponse struct {
	model.CommonResponse
	TmallSerivcecenterWorkcardInsuranceClaimAPIResponseModel
}

// Reset 清空结构体
func (m *TmallSerivcecenterWorkcardInsuranceClaimAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallSerivcecenterWorkcardInsuranceClaimAPIResponseModel).Reset()
}

// TmallSerivcecenterWorkcardInsuranceClaimAPIResponseModel is 保险理赔回传工单记录 成功返回结果
type TmallSerivcecenterWorkcardInsuranceClaimAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_serivcecenter_workcard_insurance_claim_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallSerivcecenterWorkcardInsuranceClaimAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.IsSuccess = false
}

var poolTmallSerivcecenterWorkcardInsuranceClaimAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallSerivcecenterWorkcardInsuranceClaimAPIResponse)
	},
}

// GetTmallSerivcecenterWorkcardInsuranceClaimAPIResponse 从 sync.Pool 获取 TmallSerivcecenterWorkcardInsuranceClaimAPIResponse
func GetTmallSerivcecenterWorkcardInsuranceClaimAPIResponse() *TmallSerivcecenterWorkcardInsuranceClaimAPIResponse {
	return poolTmallSerivcecenterWorkcardInsuranceClaimAPIResponse.Get().(*TmallSerivcecenterWorkcardInsuranceClaimAPIResponse)
}

// ReleaseTmallSerivcecenterWorkcardInsuranceClaimAPIResponse 将 TmallSerivcecenterWorkcardInsuranceClaimAPIResponse 保存到 sync.Pool
func ReleaseTmallSerivcecenterWorkcardInsuranceClaimAPIResponse(v *TmallSerivcecenterWorkcardInsuranceClaimAPIResponse) {
	v.Reset()
	poolTmallSerivcecenterWorkcardInsuranceClaimAPIResponse.Put(v)
}
