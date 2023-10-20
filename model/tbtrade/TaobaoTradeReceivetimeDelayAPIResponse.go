package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeReceivetimeDelayAPIResponse 延长交易收货时间 API返回值
// taobao.trade.receivetime.delay
//
// 延长交易收货时间
type TaobaoTradeReceivetimeDelayAPIResponse struct {
	model.CommonResponse
	TaobaoTradeReceivetimeDelayAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradeReceivetimeDelayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradeReceivetimeDelayAPIResponseModel).Reset()
}

// TaobaoTradeReceivetimeDelayAPIResponseModel is 延长交易收货时间 成功返回结果
type TaobaoTradeReceivetimeDelayAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_receivetime_delay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 更新后的交易数据，只包括tid和modified两个字段。
	Trade *Trade `json:"trade,omitempty" xml:"trade,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradeReceivetimeDelayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Trade = nil
}

var poolTaobaoTradeReceivetimeDelayAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradeReceivetimeDelayAPIResponse)
	},
}

// GetTaobaoTradeReceivetimeDelayAPIResponse 从 sync.Pool 获取 TaobaoTradeReceivetimeDelayAPIResponse
func GetTaobaoTradeReceivetimeDelayAPIResponse() *TaobaoTradeReceivetimeDelayAPIResponse {
	return poolTaobaoTradeReceivetimeDelayAPIResponse.Get().(*TaobaoTradeReceivetimeDelayAPIResponse)
}

// ReleaseTaobaoTradeReceivetimeDelayAPIResponse 将 TaobaoTradeReceivetimeDelayAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradeReceivetimeDelayAPIResponse(v *TaobaoTradeReceivetimeDelayAPIResponse) {
	v.Reset()
	poolTaobaoTradeReceivetimeDelayAPIResponse.Put(v)
}
