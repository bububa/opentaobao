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
    // Response *TaobaoTvpayOrderPrecreateResponse `json:"tvpay_order_precreate_response,omitempty"` 
    TaobaoTvpayOrderPrecreateResponse
}

/* model for simplify = false
type TaobaoTvpayOrderPrecreateResponse struct {

    // Top返回对象
    
    Result  *struct {
        TopResultDo  *TopResultDo `json:"top_result_do,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoTvpayOrderPrecreateResponse struct {

    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty"`

}
