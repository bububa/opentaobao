package nlife

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// Alibabanlifeb2ctraderefundAPIResponse 零售+请求退款 API返回值
// alibaba.nlife.b2c.trade.refund
//
// 零售+平台请求退款接口，在零售+平台不会有资金流变动，只是订单状态的更新
type Alibabanlifeb2ctraderefundAPIResponse struct {
	model.CommonResponse
	Alibabanlifeb2ctraderefundAPIResponseModel
}

// Alibabanlifeb2ctraderefundAPIResponseModel is 零售+请求退款 成功返回结果
type Alibabanlifeb2ctraderefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nlife_b2c_trade_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退款时间
	GmtRefund string `json:"gmt_refund,omitempty" xml:"gmt_refund,omitempty"`
	// 扩展参数
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
}
