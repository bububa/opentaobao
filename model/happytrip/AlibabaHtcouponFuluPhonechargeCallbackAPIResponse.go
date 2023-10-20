package happytrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHtcouponFuluPhonechargeCallbackAPIResponse 话费充值回调 API返回值
// alibaba.htcoupon.fulu.phonecharge.callback
//
// 话费充值为异步操作，此接口用于充值成功后，供应商回调。
type AlibabaHtcouponFuluPhonechargeCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaHtcouponFuluPhonechargeCallbackAPIResponseModel
}

// AlibabaHtcouponFuluPhonechargeCallbackAPIResponseModel is 话费充值回调 成功返回结果
type AlibabaHtcouponFuluPhonechargeCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_htcoupon_fulu_phonecharge_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回码
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// 返回码描述
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
}
