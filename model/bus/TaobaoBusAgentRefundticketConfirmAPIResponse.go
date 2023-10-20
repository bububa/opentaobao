package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusAgentRefundticketConfirmAPIResponse 商家top回调退款明细 API返回值
// taobao.bus.agent.refundticket.confirm
//
// 商家通过top回调告知平台退款明细
type TaobaoBusAgentRefundticketConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoBusAgentRefundticketConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusAgentRefundticketConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusAgentRefundticketConfirmAPIResponseModel).Reset()
}

// TaobaoBusAgentRefundticketConfirmAPIResponseModel is 商家top回调退款明细 成功返回结果
type TaobaoBusAgentRefundticketConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_agent_refundticket_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 退款回调是否收到
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusAgentRefundticketConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.IsSuccess = false
}

var poolTaobaoBusAgentRefundticketConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusAgentRefundticketConfirmAPIResponse)
	},
}

// GetTaobaoBusAgentRefundticketConfirmAPIResponse 从 sync.Pool 获取 TaobaoBusAgentRefundticketConfirmAPIResponse
func GetTaobaoBusAgentRefundticketConfirmAPIResponse() *TaobaoBusAgentRefundticketConfirmAPIResponse {
	return poolTaobaoBusAgentRefundticketConfirmAPIResponse.Get().(*TaobaoBusAgentRefundticketConfirmAPIResponse)
}

// ReleaseTaobaoBusAgentRefundticketConfirmAPIResponse 将 TaobaoBusAgentRefundticketConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusAgentRefundticketConfirmAPIResponse(v *TaobaoBusAgentRefundticketConfirmAPIResponse) {
	v.Reset()
	poolTaobaoBusAgentRefundticketConfirmAPIResponse.Put(v)
}
