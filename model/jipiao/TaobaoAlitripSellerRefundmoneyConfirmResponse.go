package jipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商订单】确认退款 APIResponse
taobao.alitrip.seller.refundmoney.confirm

代理人确认退票申请单的退款
*/
type TaobaoAlitripSellerRefundmoneyConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripSellerRefundmoneyConfirmResponse
}

type TaobaoAlitripSellerRefundmoneyConfirmResponse struct {
    XMLName xml.Name `xml:"alitrip_seller_refundmoney_confirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功确认
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
