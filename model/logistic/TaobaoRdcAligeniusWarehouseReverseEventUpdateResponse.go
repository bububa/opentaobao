package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销退单事件回传接口 APIResponse
taobao.rdc.aligenius.warehouse.reverse.event.update

用于erp回传销退单相关信息到平台
*/
type TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRdcAligeniusWarehouseReverseEventUpdateResponse `json:"rdc_aligenius_warehouse_reverse_event_update_response,omitempty"` 
    TaobaoRdcAligeniusWarehouseReverseEventUpdateResponse
}

/* model for simplify = false
type TaobaoRdcAligeniusWarehouseReverseEventUpdateResponse struct {

    // 接口返回model
    
    Result  *struct {
        TaobaoRdcAligeniusWarehouseReverseEventUpdateResult  *TaobaoRdcAligeniusWarehouseReverseEventUpdateResult `json:"taobao_rdc_aligenius_warehouse_reverse_event_update_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoRdcAligeniusWarehouseReverseEventUpdateResponse struct {

    // 接口返回model
    Result   *TaobaoRdcAligeniusWarehouseReverseEventUpdateResult `json:"result,omitempty"`

}
