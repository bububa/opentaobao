package jipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
机票代理商】回填手续费 APIResponse
taobao.alitrip.seller.refund.fillfee

回填手续费
*/
type TaobaoAlitripSellerRefundFillfeeAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripSellerRefundFillfeeResponse
}

type TaobaoAlitripSellerRefundFillfeeResponse struct {
    XMLName xml.Name `xml:"alitrip_seller_refund_fillfee_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
