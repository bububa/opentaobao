package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxywechatpaycallbackAPIResponse 微信支付回调 API返回值
// alitrip.merchant.galaxy.wechat.pay.callback
//
// 微信支付回调
type AlitripmerchantgalaxywechatpaycallbackAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxywechatpaycallbackAPIResponseModel
}

// AlitripmerchantgalaxywechatpaycallbackAPIResponseModel is 微信支付回调 成功返回结果
type AlitripmerchantgalaxywechatpaycallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_pay_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *AlitripmerchantgalaxywechatpaycallbackResponse `json:"result,omitempty" xml:"result,omitempty"`
}
