package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusAgentBookticketConfirmAPIResponse 汽车票代理商接口—确认出票是否成功 API返回值
// taobao.bus.agent.bookticket.confirm
//
// 代理商通过该接口通知汽车票系统订单出票结果。
type TaobaoBusAgentBookticketConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoBusAgentBookticketConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusAgentBookticketConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusAgentBookticketConfirmAPIResponseModel).Reset()
}

// TaobaoBusAgentBookticketConfirmAPIResponseModel is 汽车票代理商接口—确认出票是否成功 成功返回结果
type TaobaoBusAgentBookticketConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_agent_bookticket_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 是否确认成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusAgentBookticketConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.IsSuccess = false
}

var poolTaobaoBusAgentBookticketConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusAgentBookticketConfirmAPIResponse)
	},
}

// GetTaobaoBusAgentBookticketConfirmAPIResponse 从 sync.Pool 获取 TaobaoBusAgentBookticketConfirmAPIResponse
func GetTaobaoBusAgentBookticketConfirmAPIResponse() *TaobaoBusAgentBookticketConfirmAPIResponse {
	return poolTaobaoBusAgentBookticketConfirmAPIResponse.Get().(*TaobaoBusAgentBookticketConfirmAPIResponse)
}

// ReleaseTaobaoBusAgentBookticketConfirmAPIResponse 将 TaobaoBusAgentBookticketConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusAgentBookticketConfirmAPIResponse(v *TaobaoBusAgentBookticketConfirmAPIResponse) {
	v.Reset()
	poolTaobaoBusAgentBookticketConfirmAPIResponse.Put(v)
}
