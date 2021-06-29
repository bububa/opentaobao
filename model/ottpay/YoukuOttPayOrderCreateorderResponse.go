package ottpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建订单 APIResponse
youku.ott.pay.order.createorder

ottpay创建订单
*/
type YoukuOttPayOrderCreateorderAPIResponse struct {
    model.CommonResponse
    YoukuOttPayOrderCreateorderResponse
}

type YoukuOttPayOrderCreateorderResponse struct {
    XMLName xml.Name `xml:"youku_ott_pay_order_createorder_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // data
    
    Data   *TvOrderResultDTO `json:"data,omitempty" xml:"data,omitempty"`

    
}
