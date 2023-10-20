package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusAgentRefundConfirmAPIResponse 汽车票退票和退款二合一接口 API返回值
// taobao.bus.agent.refund.confirm
//
// 1.商家退票成功后，回调飞猪平台汽车票退票接口，平台进行退票和退款操作。
type TaobaoBusAgentRefundConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoBusAgentRefundConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusAgentRefundConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusAgentRefundConfirmAPIResponseModel).Reset()
}

// TaobaoBusAgentRefundConfirmAPIResponseModel is 汽车票退票和退款二合一接口 成功返回结果
type TaobaoBusAgentRefundConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_agent_refund_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 退票回调是否收到
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusAgentRefundConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.IsSuccess = false
}

var poolTaobaoBusAgentRefundConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusAgentRefundConfirmAPIResponse)
	},
}

// GetTaobaoBusAgentRefundConfirmAPIResponse 从 sync.Pool 获取 TaobaoBusAgentRefundConfirmAPIResponse
func GetTaobaoBusAgentRefundConfirmAPIResponse() *TaobaoBusAgentRefundConfirmAPIResponse {
	return poolTaobaoBusAgentRefundConfirmAPIResponse.Get().(*TaobaoBusAgentRefundConfirmAPIResponse)
}

// ReleaseTaobaoBusAgentRefundConfirmAPIResponse 将 TaobaoBusAgentRefundConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusAgentRefundConfirmAPIResponse(v *TaobaoBusAgentRefundConfirmAPIResponse) {
	v.Reset()
	poolTaobaoBusAgentRefundConfirmAPIResponse.Put(v)
}
