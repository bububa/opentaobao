package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjOcPayAPIResponse
POS收银成功后订单同步 API返回值
alibaba.mj.oc.pay

此API用于在银泰商场中，消费者在收银台收银/退款时， POS系统在收银或退款成功后，调用此接口进行订单同步 */
type AlibabaMjOcPayAPIResponse struct {
	model.CommonResponse
	AlibabaMjOcPayAPIResponseModel
}

// AlibabaMjOcPayAPIResponseModel is POS收银成功后订单同步 成功返回结果
type AlibabaMjOcPayAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_oc_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errCode
	ExCode int64 `json:"ex_code,omitempty" xml:"ex_code,omitempty"`
	// errMsg
	ExMsg string `json:"ex_msg,omitempty" xml:"ex_msg,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// POS交易号
	OutTradeNo string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`
}
