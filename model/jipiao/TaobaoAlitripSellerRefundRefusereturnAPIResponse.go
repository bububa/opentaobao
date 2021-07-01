package jipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商】拒绝退票 API返回值 
taobao.alitrip.seller.refund.refusereturn

拒绝退票
*/
type TaobaoAlitripSellerRefundRefusereturnAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripSellerRefundRefusereturnAPIResponseModel
}

// 【机票代理商】拒绝退票 成功返回结果
type TaobaoAlitripSellerRefundRefusereturnAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_seller_refund_refusereturn_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *TaobaoAlitripSellerRefundRefusereturnResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
