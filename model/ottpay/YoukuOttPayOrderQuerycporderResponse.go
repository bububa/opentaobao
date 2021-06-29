package ottpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询支付订单对应cp订单号 APIResponse
youku.ott.pay.order.querycporder

根据支付订单查询对应cp订单号
*/
type YoukuOttPayOrderQuerycporderAPIResponse struct {
    model.CommonResponse
    YoukuOttPayOrderQuerycporderResponse
}

type YoukuOttPayOrderQuerycporderResponse struct {
    XMLName xml.Name `xml:"youku_ott_pay_order_querycporder_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // data
    
    Data   *TvOrderResultDTO `json:"data,omitempty" xml:"data,omitempty"`

    
}
