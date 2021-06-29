package jipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商】拒绝退票 APIResponse
taobao.alitrip.seller.refund.refusereturn

拒绝退票
*/
type TaobaoAlitripSellerRefundRefusereturnAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripSellerRefundRefusereturnResponse
}

type TaobaoAlitripSellerRefundRefusereturnResponse struct {
    XMLName xml.Name `xml:"alitrip_seller_refund_refusereturn_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TaobaoAlitripSellerRefundRefusereturnResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
