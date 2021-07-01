package openmall

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallTradeAgreepayAPIResponse
openmall订单支付 API返回值
taobao.openmall.trade.agreepay

openmall订单支付 */
type TaobaoOpenmallTradeAgreepayAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallTradeAgreepayAPIResponseModel
}

// TaobaoOpenmallTradeAgreepayAPIResponseModel is openmall订单支付 成功返回结果
type TaobaoOpenmallTradeAgreepayAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_trade_agreepay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
