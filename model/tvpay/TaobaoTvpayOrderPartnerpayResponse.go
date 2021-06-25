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
    Response *TaobaoTvpayOrderPartnerpayResponse `json:"taobao_tvpay_order_partnerpay_response,omitempty"`
}

type TaobaoTvpayOrderPartnerpayResponse struct {

    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty"`

}
