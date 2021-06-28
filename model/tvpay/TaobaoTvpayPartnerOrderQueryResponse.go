package tvpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商户查询订单 APIResponse
taobao.tvpay.partner.order.query

给商户提供的查询订单状态的API
*/
type TaobaoTvpayPartnerOrderQueryAPIResponse struct {
    model.CommonResponse
    TaobaoTvpayPartnerOrderQueryResponse
}

type TaobaoTvpayPartnerOrderQueryResponse struct {
    XMLName xml.Name `xml:"tvpay_partner_order_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // Top返回对象
    
    Result   *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
