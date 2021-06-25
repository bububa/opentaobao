package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取销退收货信息 APIResponse
taobao.wlb.wms.return.bill.get

通过订单号获取单个销退单收货信息。
*/
type TaobaoWlbWmsReturnBillGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbWmsReturnBillGetResponse `json:"taobao_wlb_wms_return_bill_get_response,omitempty"`
}

type TaobaoWlbWmsReturnBillGetResponse struct {

    // 回退订单信息
    ReturnOrderInfo   *CainiaoReturnBillReturnorderinfo `json:"return_order_info,omitempty"`

}
