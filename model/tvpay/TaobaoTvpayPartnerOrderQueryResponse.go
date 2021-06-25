package tvpay

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商户查询订单 APIResponse
taobao.tvpay.partner.order.query

给商户提供的查询订单状态的API
*/
type TaobaoTvpayPartnerOrderQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTvpayPartnerOrderQueryResponse `json:"taobao_tvpay_partner_order_query_response,omitempty"`
}

type TaobaoTvpayPartnerOrderQueryResponse struct {

    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty"`

}
