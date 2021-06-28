package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取销售订单发货信息 APIResponse
taobao.wlb.wms.consign.bill.get

获取销售订单发货信息
*/
type TaobaoWlbWmsConsignBillGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbWmsConsignBillGetResponse `json:"wlb_wms_consign_bill_get_response,omitempty"` 
    TaobaoWlbWmsConsignBillGetResponse
}

/* model for simplify = false
type TaobaoWlbWmsConsignBillGetResponse struct {

    // 商品信息列表
    
    ConsignSendInfoList  struct {
        Consignsendinfolist  []Consignsendinfolist `json:"consignsendinfolist,omitempty"`
    } `json:"consign_send_info_list,omitempty"`
    

}
*/

type TaobaoWlbWmsConsignBillGetResponse struct {

    // 商品信息列表
    ConsignSendInfoList   []Consignsendinfolist `json:"consign_send_info_list,omitempty"`

}
