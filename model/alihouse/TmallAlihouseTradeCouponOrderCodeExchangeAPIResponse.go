package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallAlihouseTradeCouponOrderCodeExchangeAPIResponse 核销券码 API返回值
// tmall.alihouse.trade.coupon.order.code.exchange
//
// ETC核销券码
type TmallAlihouseTradeCouponOrderCodeExchangeAPIResponse struct {
	model.CommonResponse
	TmallAlihouseTradeCouponOrderCodeExchangeAPIResponseModel
}

// TmallAlihouseTradeCouponOrderCodeExchangeAPIResponseModel is 核销券码 成功返回结果
type TmallAlihouseTradeCouponOrderCodeExchangeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_alihouse_trade_coupon_order_code_exchange_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallAlihouseTradeCouponOrderCodeExchangeResult `json:"result,omitempty" xml:"result,omitempty"`
}
