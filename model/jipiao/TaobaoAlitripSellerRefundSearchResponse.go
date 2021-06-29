package jipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商】退票申请单列表 APIResponse
taobao.alitrip.seller.refund.search

查询退票申请单列表
*/
type TaobaoAlitripSellerRefundSearchAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripSellerRefundSearchResponse
}

type TaobaoAlitripSellerRefundSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_seller_refund_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TaobaoAlitripSellerRefundSearchResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
