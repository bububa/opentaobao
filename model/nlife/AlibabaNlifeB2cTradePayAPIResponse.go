package nlife

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cTradePayAPIResponse 零售+平台支付订单 API返回值
// alibaba.nlife.b2c.trade.pay
//
// 零售+平台支付接口，外部商户调用此接口告知零售+支付结果，保持订单状态同步
type AlibabaNlifeB2cTradePayAPIResponse struct {
	model.CommonResponse
	AlibabaNlifeB2cTradePayAPIResponseModel
}

// AlibabaNlifeB2cTradePayAPIResponseModel is 零售+平台支付订单 成功返回结果
type AlibabaNlifeB2cTradePayAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nlife_b2c_trade_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// gmtPayment
	GmtPayment string `json:"gmt_payment,omitempty" xml:"gmt_payment,omitempty"`
	// 扩展参数
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
}
