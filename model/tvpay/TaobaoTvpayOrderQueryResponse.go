package tvpay

import (
    "github.com/bububa/opentaobao/model"
)

/* 
tv支付查询订单状态 APIResponse
taobao.tvpay.order.query

tv支付查询订单状态
*/
type TaobaoTvpayOrderQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTvpayOrderQueryResponse `json:"taobao_tvpay_order_query_response,omitempty"`
}

type TaobaoTvpayOrderQueryResponse struct {

    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty"`

}
