package jipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
敏感信息查询 API返回值 
taobao.alitrip.buyer.get

针对商家提供统一的TOP接口，可以根据订单获取订单对应买家联系电话（阿里小号）。
*/
type TaobaoAlitripBuyerGetAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripBuyerGetResponse
}

// 敏感信息查询 成功返回结果
type TaobaoAlitripBuyerGetResponse struct {
    XMLName xml.Name `xml:"alitrip_buyer_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 请求内容，如阿里小号
    Content   string `json:"content,omitempty" xml:"content,omitempty"`
    // 有效期
    Expires   string `json:"expires,omitempty" xml:"expires,omitempty"`
}
