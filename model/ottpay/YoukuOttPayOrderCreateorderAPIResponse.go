package ottpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建订单 API返回值 
youku.ott.pay.order.createorder

ottpay创建订单
*/
type YoukuOttPayOrderCreateorderAPIResponse struct {
    model.CommonResponse
    YoukuOttPayOrderCreateorderAPIResponseModel
}

// 创建订单 成功返回结果
type YoukuOttPayOrderCreateorderAPIResponseModel struct {
    XMLName xml.Name `xml:"youku_ott_pay_order_createorder_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // data
    Data   *TvOrderResultDto `json:"data,omitempty" xml:"data,omitempty"`
}
