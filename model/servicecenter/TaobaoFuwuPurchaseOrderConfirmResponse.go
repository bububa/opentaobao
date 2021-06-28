package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务市场内购服务下单接口 APIResponse
taobao.fuwu.purchase.order.confirm

通过传入服务市场商品的itemcode等信息，返回给服务商内购服务的下单链接
*/
type TaobaoFuwuPurchaseOrderConfirmAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"fuwu_purchase_order_confirm_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 下单页面url
    
    Url   string `json:"url,omitempty" xml:"