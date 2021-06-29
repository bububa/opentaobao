package ottpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
委托代扣服务 APIResponse
youku.ott.pay.order.authpay

应用中心sdk连续包月委托代扣服务
*/
type YoukuOttPayOrderAuthpayAPIResponse struct {
    model.CommonResponse
    YoukuOttPayOrderAuthpayResponse
}

type YoukuOttPayOrderAuthpayResponse struct {
    XMLName xml.Name `xml:"youku_ott_pay_order_authpay_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // data
    
    Data   *TvOrderResultDto `json:"data,omitempty" xml:"data,omitempty"`

    
}
