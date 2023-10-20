package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketFlowConsumeAPIResponse 无交易类凭证核销 API返回值
// taobao.vmarket.eticket.flow.consume
//
// 无交易类凭证核销
type TaobaoVmarketEticketFlowConsumeAPIResponse struct {
	model.CommonResponse
	TaobaoVmarketEticketFlowConsumeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketFlowConsumeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoVmarketEticketFlowConsumeAPIResponseModel).Reset()
}

// TaobaoVmarketEticketFlowConsumeAPIResponseModel is 无交易类凭证核销 成功返回结果
type TaobaoVmarketEticketFlowConsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_flow_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误提示信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 执行成功
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketFlowConsumeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.RetCode = 0
}

var poolTaobaoVmarketEticketFlowConsumeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoVmarketEticketFlowConsumeAPIResponse)
	},
}

// GetTaobaoVmarketEticketFlowConsumeAPIResponse 从 sync.Pool 获取 TaobaoVmarketEticketFlowConsumeAPIResponse
func GetTaobaoVmarketEticketFlowConsumeAPIResponse() *TaobaoVmarketEticketFlowConsumeAPIResponse {
	return poolTaobaoVmarketEticketFlowConsumeAPIResponse.Get().(*TaobaoVmarketEticketFlowConsumeAPIResponse)
}

// ReleaseTaobaoVmarketEticketFlowConsumeAPIResponse 将 TaobaoVmarketEticketFlowConsumeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoVmarketEticketFlowConsumeAPIResponse(v *TaobaoVmarketEticketFlowConsumeAPIResponse) {
	v.Reset()
	poolTaobaoVmarketEticketFlowConsumeAPIResponse.Put(v)
}
