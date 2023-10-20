package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousemerchanttradeconfigbindAPIResponse 交易场景绑定 API返回值
// alibaba.alihouse.merchant.trade.config.bind
//
// 交易场景绑定
type AlibabaalihousemerchanttradeconfigbindAPIResponse struct {
	model.CommonResponse
	AlibabaalihousemerchanttradeconfigbindAPIResponseModel
}

// AlibabaalihousemerchanttradeconfigbindAPIResponseModel is 交易场景绑定 成功返回结果
type AlibabaalihousemerchanttradeconfigbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_merchant_trade_config_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousemerchanttradeconfigbindResult `json:"result,omitempty" xml:"result,omitempty"`
}
