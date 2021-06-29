package ottpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单查询接口(cp订单号查询) APIResponse
youku.ott.pay.order.queryorderbycp

给商户服务端查询订单状态
*/
type YoukuOttPayOrderQueryorderbycpAPIResponse struct {
    model.CommonResponse
    YoukuOttPayOrderQueryorderbycpResponse
}

type YoukuOttPayOrderQueryorderbycpResponse struct {
    XMLName xml.Name `xml:"youku_ott_pay_order_queryorderbycp_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // data
    
    Data   *TvOrderQueryResultDTO `json:"data,omitempty" xml:"data,omitempty"`

    
}
