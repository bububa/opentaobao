package jipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商】确认退票 API返回值 
taobao.alitrip.seller.refund.confirmreturn

确认退票
*/
type TaobaoAlitripSellerRefundConfirmreturnAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripSellerRefundConfirmreturnResponse
}

// 【机票代理商】确认退票 成功返回结果
type TaobaoAlitripSellerRefundConfirmreturnResponse struct {
    XMLName xml.Name `xml:"alitrip_seller_refund_confirmreturn_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否成功
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
