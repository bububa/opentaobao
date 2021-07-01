package ottpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询支付订单对应cp订单号 API返回值 
youku.ott.pay.order.querycporder

根据支付订单查询对应cp订单号
*/
type YoukuOttPayOrderQuerycporderAPIResponse struct {
    model.CommonResponse
    YoukuOttPayOrderQuerycporderAPIResponseModel
}

// 查询支付订单对应cp订单号 成功返回结果
type YoukuOttPayOrderQuerycporderAPIResponseModel struct {
    XMLName xml.Name `xml:"youku_ott_pay_order_querycporder_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // data
    Data   *TvOrderResultDto `json:"data,omitempty" xml:"data,omitempty"`
}
