package nlife

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeB2cTradeCancelAPIResponse
零售+平台取消订单 API返回值
alibaba.nlife.b2c.trade.cancel

零售+平台取消订单接口 */
type AlibabaNlifeB2cTradeCancelAPIResponse struct {
	model.CommonResponse
	AlibabaNlifeB2cTradeCancelAPIResponseModel
}

// AlibabaNlifeB2cTradeCancelAPIResponseModel is 零售+平台取消订单 成功返回结果
type AlibabaNlifeB2cTradeCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nlife_b2c_trade_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单取消时间
	GmtCancel string `json:"gmt_cancel,omitempty" xml:"gmt_cancel,omitempty"`
	// 扩展参数JSON
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
}
