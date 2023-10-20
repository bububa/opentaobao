package eticket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketResendAPIResponse 外部合作商家重发电子凭证回调接口 API返回值
// taobao.vmarket.eticket.resend
//
// 外部合作商家重发电子凭证回调接口
type TaobaoVmarketEticketResendAPIResponse struct {
	model.CommonResponse
	TaobaoVmarketEticketResendAPIResponseModel
}

// TaobaoVmarketEticketResendAPIResponseModel is 外部合作商家重发电子凭证回调接口 成功返回结果
type TaobaoVmarketEticketResendAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_resend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 0:失败，1:成功
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}
