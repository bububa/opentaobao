package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
打印确认接口v1.0 APIResponse
taobao.wlb.waybill.i.print

打印面单前的校验接口，判断面单号信息与订单信息是否匹配。
*/
type TaobaoWlbWaybillIPrintAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbWaybillIPrintResponse `json:"wlb_waybill_i_print_response,omitempty"` 
    TaobaoWlbWaybillIPrintResponse
}

/* model for simplify = false
type TaobaoWlbWaybillIPrintResponse struct {

    // 面单打印信息
    
    WaybillApplyPrintCheckInfos  struct {
        WaybillApplyPrintCheckInfo  []WaybillApplyPrintCheckInfo `json:"waybill_apply_print_check_info,omitempty"`
    } `json:"waybill_apply_print_check_infos,omitempty"`
    

}
*/

type TaobaoWlbWaybillIPrintResponse struct {

    // 面单打印信息
    WaybillApplyPrintCheckInfos   []WaybillApplyPrintCheckInfo `json:"waybill_apply_print_check_infos,omitempty"`

}
