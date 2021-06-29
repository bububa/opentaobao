package ottpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退订应用中心支付订单 API返回值 
youku.ott.pay.order.deleteorder

应用中心sdk连续包月退订接口
*/
type YoukuOttPayOrderDeleteorderAPIResponse struct {
    model.CommonResponse
    YoukuOttPayOrderDeleteorderResponse
}

// 退订应用中心支付订单 成功返回结果
type YoukuOttPayOrderDeleteorderResponse struct {
    XMLName xml.Name `xml:"youku_ott_pay_order_deleteorder_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // data
    Data   *TvOrderResultDTO `json:"data,omitempty" xml:"data,omitempty"`
}
