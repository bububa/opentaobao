package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketFlowResendAPIResponse 业务重新触发发码短信 API返回值
// taobao.vmarket.eticket.flow.resend
//
// 业务重新触发发码短信
type TaobaoVmarketEticketFlowResendAPIResponse struct {
	model.CommonResponse
	TaobaoVmarketEticketFlowResendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketFlowResendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoVmarketEticketFlowResendAPIResponseModel).Reset()
}

// TaobaoVmarketEticketFlowResendAPIResponseModel is 业务重新触发发码短信 成功返回结果
type TaobaoVmarketEticketFlowResendAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_flow_resend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误提示信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 1成功;0失败
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketFlowResendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.RetCode = 0
}

var poolTaobaoVmarketEticketFlowResendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoVmarketEticketFlowResendAPIResponse)
	},
}

// GetTaobaoVmarketEticketFlowResendAPIResponse 从 sync.Pool 获取 TaobaoVmarketEticketFlowResendAPIResponse
func GetTaobaoVmarketEticketFlowResendAPIResponse() *TaobaoVmarketEticketFlowResendAPIResponse {
	return poolTaobaoVmarketEticketFlowResendAPIResponse.Get().(*TaobaoVmarketEticketFlowResendAPIResponse)
}

// ReleaseTaobaoVmarketEticketFlowResendAPIResponse 将 TaobaoVmarketEticketFlowResendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoVmarketEticketFlowResendAPIResponse(v *TaobaoVmarketEticketFlowResendAPIResponse) {
	v.Reset()
	poolTaobaoVmarketEticketFlowResendAPIResponse.Put(v)
}
