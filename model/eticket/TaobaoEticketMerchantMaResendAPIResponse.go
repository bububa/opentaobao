package eticket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoEticketMerchantMaResendAPIResponse 电子凭证重发回调接口 API返回值
// taobao.eticket.merchant.ma.resend
//
// 码商重发电子凭证回调接口
type TaobaoEticketMerchantMaResendAPIResponse struct {
	model.CommonResponse
	TaobaoEticketMerchantMaResendAPIResponseModel
}

// TaobaoEticketMerchantMaResendAPIResponseModel is 电子凭证重发回调接口 成功返回结果
type TaobaoEticketMerchantMaResendAPIResponseModel struct {
	XMLName xml.Name `xml:"eticket_merchant_ma_resend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回复参数
	RespBody *SendMaCallbackResp `json:"resp_body,omitempty" xml:"resp_body,omitempty"`
	// 子结果码
	RetCode string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 子结果信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
}
