package ottpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退订应用中心支付订单 APIResponse
youku.ott.pay.order.deleteorder

应用中心sdk连续包月退订接口
*/
type YoukuOttPayOrderDeleteorderAPIResponse struct {
    model.CommonResponse
    YoukuOttPayOrderDeleteorderResponse
}

type YoukuOttPayOrderDeleteorderResponse struct {
    XMLName xml.Name `xml:"youku_ott_pay_order_deleteorder_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // data
    
    Data   *TvOrderResultDto `json:"data,omitempty" xml:"data,omitempty"`

    
}
