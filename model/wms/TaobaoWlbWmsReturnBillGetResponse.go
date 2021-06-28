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
    // Response *TaobaoWlbWmsReturnBillGetResponse `json:"wlb_wms_return_bill_get_response,omitempty"` 
    TaobaoWlbWmsReturnBillGetResponse
}

/* model for simplify = false
type TaobaoWlbWmsReturnBillGetResponse struct {

    // 回退订单信息
    
    ReturnOrderInfo  *struct {
        CainiaoReturnBillReturnorderinfo  *CainiaoReturnBillReturnorderinfo `json:"cainiao_return_bill_returnorderinfo,omitempty"`
    } `json:"return_order_info,omitempty"`
    

}
*/

type TaobaoWlbWmsReturnBillGetResponse struct {

    // 回退订单信息
    ReturnOrderInfo   *CainiaoReturnBillReturnorderinfo `json:"return_order_info,omitempty"`

}
