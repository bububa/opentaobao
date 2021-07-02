package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWtOrderExchangePartnerChecktbuserAPIResponse 积分兑换校验淘宝账号是否存在 API返回值
// alibaba.wt.order.exchange.partner.checktbuser
//
// 积分兑换之前校验淘宝账号是否存在
type AlibabaWtOrderExchangePartnerChecktbuserAPIResponse struct {
	model.CommonResponse
	AlibabaWtOrderExchangePartnerChecktbuserAPIResponseModel
}

// AlibabaWtOrderExchangePartnerChecktbuserAPIResponseModel is 积分兑换校验淘宝账号是否存在 成功返回结果
type AlibabaWtOrderExchangePartnerChecktbuserAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wt_order_exchange_partner_checktbuser_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值，通过model的值true或者false来判断
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 成功
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 成功
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 接口调用返回成功，真正是否存在号码通过model的返回值来判断
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
