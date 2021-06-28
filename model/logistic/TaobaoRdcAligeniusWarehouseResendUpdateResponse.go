package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
补发单状态回传 APIResponse
taobao.rdc.aligenius.warehouse.resend.update

补发单状态回传接口
*/
type TaobaoRdcAligeniusWarehouseResendUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRdcAligeniusWarehouseResendUpdateResponse `json:"rdc_aligenius_warehouse_resend_update_response,omitempty"` 
    TaobaoRdcAligeniusWarehouseResendUpdateResponse
}

/* model for simplify = false
type TaobaoRdcAligeniusWarehouseResendUpdateResponse struct {

    // result
    
    Result  *struct {
        TaobaoRdcAligeniusWarehouseResendUpdateResult  *TaobaoRdcAligeniusWarehouseResendUpdateResult `json:"taobao_rdc_aligenius_warehouse_resend_update_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoRdcAligeniusWarehouseResendUpdateResponse struct {

    // result
    Result   *TaobaoRdcAligeniusWarehouseResendUpdateResult `json:"result,omitempty"`

}
