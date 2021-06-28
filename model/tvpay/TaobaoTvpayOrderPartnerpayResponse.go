package tvpay

import (
    "github.com/bububa/opentaobao/model"
)

/* 
tv支付第三方支付订单 APIResponse
taobao.tvpay.order.partnerpay

tv支付第三方发起并支付订单（使用设备授权）
*/
type TaobaoTvpayOrderPartnerpayAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTvpayOrderPartnerpayResponse `json:"tvpay_order_partnerpay_response,omitempty"` 
    TaobaoTvpayOrderPartnerpayResponse
}

/* model for simplify = false
type TaobaoTvpayOrderPartnerpayResponse struct {

    // Top返回对象
    
    Result  *struct {
        TopResultDo  *TopResultDo `json:"top_result_do,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoTvpayOrderPartnerpayResponse struct {

    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty"`

}
