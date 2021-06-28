package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取物流服务商电子面单号v1.0 APIResponse
taobao.wlb.waybill.i.get

商家根据订单信息，实时、批量获取指定物流服务商的电子面单号。
*/
type TaobaoWlbWaybillIGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbWaybillIGetResponse `json:"wlb_waybill_i_get_response,omitempty"` 
    TaobaoWlbWaybillIGetResponse
}

/* model for simplify = false
type TaobaoWlbWaybillIGetResponse struct {

    // 面单申请接口返回信息
    
    WaybillApplyNewCols  struct {
        WaybillApplyNewInfo  []WaybillApplyNewInfo `json:"waybill_apply_new_info,omitempty"`
    } `json:"waybill_apply_new_cols,omitempty"`
    

}
*/

type TaobaoWlbWaybillIGetResponse struct {

    // 面单申请接口返回信息
    WaybillApplyNewCols   []WaybillApplyNewInfo `json:"waybill_apply_new_cols,omitempty"`

}
