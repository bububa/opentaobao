package tbtrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeConfirmfeeGetAPIResponse 获取交易确认收货费用 API返回值
// taobao.trade.confirmfee.get
//
// 获取交易确认收货费用，可以获取主订单或子订单的确认收货费用
type TaobaoTradeConfirmfeeGetAPIResponse struct {
	model.CommonResponse
	TaobaoTradeConfirmfeeGetAPIResponseModel
}

// TaobaoTradeConfirmfeeGetAPIResponseModel is 获取交易确认收货费用 成功返回结果
type TaobaoTradeConfirmfeeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_confirmfee_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取到的交易确认收货费用
	TradeConfirmFee *TradeConfirmFee `json:"trade_confirm_fee,omitempty" xml:"trade_confirm_fee,omitempty"`
}
