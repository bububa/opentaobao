package jipiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripsellerrefundmoneyconfirmAPIResponse 【机票代理商订单】确认退款 API返回值
// taobao.alitrip.seller.refundmoney.confirm
//
// 代理人确认退票申请单的退款
type TaobaoalitripsellerrefundmoneyconfirmAPIResponse struct {
	model.CommonResponse
	TaobaoalitripsellerrefundmoneyconfirmAPIResponseModel
}

// TaobaoalitripsellerrefundmoneyconfirmAPIResponseModel is 【机票代理商订单】确认退款 成功返回结果
type TaobaoalitripsellerrefundmoneyconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_seller_refundmoney_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功确认
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
