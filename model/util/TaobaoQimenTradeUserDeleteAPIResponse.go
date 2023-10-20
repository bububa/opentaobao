package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenTradeUserDeleteAPIResponse 删除奇门订单链路用户 API返回值
// taobao.qimen.trade.user.delete
//
// 删除奇门订单链路用户
type TaobaoQimenTradeUserDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoQimenTradeUserDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenTradeUserDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenTradeUserDeleteAPIResponseModel).Reset()
}

// TaobaoQimenTradeUserDeleteAPIResponseModel is 删除奇门订单链路用户 成功返回结果
type TaobaoQimenTradeUserDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_trade_user_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// modal
	Modal bool `json:"modal,omitempty" xml:"modal,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenTradeUserDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Modal = false
}

var poolTaobaoQimenTradeUserDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenTradeUserDeleteAPIResponse)
	},
}

// GetTaobaoQimenTradeUserDeleteAPIResponse 从 sync.Pool 获取 TaobaoQimenTradeUserDeleteAPIResponse
func GetTaobaoQimenTradeUserDeleteAPIResponse() *TaobaoQimenTradeUserDeleteAPIResponse {
	return poolTaobaoQimenTradeUserDeleteAPIResponse.Get().(*TaobaoQimenTradeUserDeleteAPIResponse)
}

// ReleaseTaobaoQimenTradeUserDeleteAPIResponse 将 TaobaoQimenTradeUserDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenTradeUserDeleteAPIResponse(v *TaobaoQimenTradeUserDeleteAPIResponse) {
	v.Reset()
	poolTaobaoQimenTradeUserDeleteAPIResponse.Put(v)
}
