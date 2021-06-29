package jipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商】退票申请单详情 APIResponse
taobao.alitrip.seller.refund.get

查询退票申请单详情
*/
type TaobaoAlitripSellerRefundGetAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripSellerRefundGetResponse
}

type TaobaoAlitripSellerRefundGetResponse struct {
    XMLName xml.Name `xml:"alitrip_seller_refund_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TaobaoAlitripSellerRefundGetResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
