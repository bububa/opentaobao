package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
内购服务订单付款页获取接口 APIResponse
taobao.fuwu.purchase.order.pay

通过接口获取某一订单的付款页面链接
*/
type TaobaoFuwuPurchaseOrderPayAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFuwuPurchaseOrderPayResponse `json:"taobao_fuwu_purchase_order_pay_response,omitempty"`
}

type TaobaoFuwuPurchaseOrderPayResponse struct {

    // 该url用于订单付款
    Url   string `json:"url,omitempty"`

}
