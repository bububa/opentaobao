package ottpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询订单 APIResponse
youku.ott.pay.order.queryorder

通过订单号查询订单信息
*/
type YoukuOttPayOrderQueryorderAPIResponse struct {
    model.CommonResponse
    YoukuOttPayOrderQueryorderResponse
}

type YoukuOttPayOrderQueryorderResponse struct {
    XMLName xml.Name `xml:"youku_ott_pay_order_queryorder_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // status
    
    Data   *TvOrderQueryResultDTO `json:"data,omitempty" xml:"data,omitempty"`

    
}
