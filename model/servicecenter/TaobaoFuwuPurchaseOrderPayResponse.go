package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
内购服务订单付款页获取接口 APIResponse
taobao.fuwu.purchase.order.pay

通过接口获取某一订单的付款页面链接
*/
type TaobaoFuwuPurchaseOrderPayAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"fuwu_purchase_order_pay_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 该url用于订单付款
    
    Url   string `json:"url,omitempty" xml:"