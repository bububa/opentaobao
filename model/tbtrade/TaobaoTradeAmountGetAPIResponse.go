package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeAmountGetAPIResponse 交易帐务查询 API返回值
// taobao.trade.amount.get
//
// 卖家查询该笔交易的资金帐务相关的数据；
// 1. 只供卖家使用，买家不可使用
// 2. 可查询所有的状态的交易，但不同状态时交易的相关数据可能会有不同
type TaobaoTradeAmountGetAPIResponse struct {
	model.CommonResponse
	TaobaoTradeAmountGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradeAmountGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradeAmountGetAPIResponseModel).Reset()
}

// TaobaoTradeAmountGetAPIResponseModel is 交易帐务查询 成功返回结果
type TaobaoTradeAmountGetAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_amount_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 主订单的财务信息详情
	TradeAmount *TradeAmount `json:"trade_amount,omitempty" xml:"trade_amount,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradeAmountGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TradeAmount = nil
}

var poolTaobaoTradeAmountGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradeAmountGetAPIResponse)
	},
}

// GetTaobaoTradeAmountGetAPIResponse 从 sync.Pool 获取 TaobaoTradeAmountGetAPIResponse
func GetTaobaoTradeAmountGetAPIResponse() *TaobaoTradeAmountGetAPIResponse {
	return poolTaobaoTradeAmountGetAPIResponse.Get().(*TaobaoTradeAmountGetAPIResponse)
}

// ReleaseTaobaoTradeAmountGetAPIResponse 将 TaobaoTradeAmountGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradeAmountGetAPIResponse(v *TaobaoTradeAmountGetAPIResponse) {
	v.Reset()
	poolTaobaoTradeAmountGetAPIResponse.Put(v)
}
