package tvpay

import (
    "github.com/bububa/opentaobao/model"
)

/* 
tv支付预下单 APIResponse
taobao.tvpay.order.precreate

tv支付预下单
*/
type TaobaoTvpayOrderPrecreateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTvpayOrderPrecreateResponse `json:"taobao_tvpay_order_precreate_response,omitempty"`
}

type TaobaoTvpayOrderPrecreateResponse struct {

    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty"`

}
