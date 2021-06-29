package jipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商】确认退票 APIResponse
taobao.alitrip.seller.refund.confirmreturn

确认退票
*/
type TaobaoAlitripSellerRefundConfirmreturnAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripSellerRefundConfirmreturnResponse
}

type TaobaoAlitripSellerRefundConfirmreturnResponse struct {
    XMLName xml.Name `xml:"alitrip_seller_refund_confirmreturn_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
