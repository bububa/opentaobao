package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosOnsiteTradeIsnewpayorderAPIResponse
是否为新支付订单 API返回值
alibaba.mos.onsite.trade.isnewpayorder

退款时，老支付宝手淘退款接口需要查一下该订单是否为新支付订单 */
type AlibabaMosOnsiteTradeIsnewpayorderAPIResponse struct {
	model.CommonResponse
	AlibabaMosOnsiteTradeIsnewpayorderAPIResponseModel
}

// AlibabaMosOnsiteTradeIsnewpayorderAPIResponseModel is 是否为新支付订单 成功返回结果
type AlibabaMosOnsiteTradeIsnewpayorderAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_onsite_trade_isnewpayorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaMosOnsiteTradeIsnewpayorderResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
