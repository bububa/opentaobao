package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
悬挂链状态回传确认 APIResponse
taobao.wdk.equipment.conveyor.statusconfirm

悬挂链状态回传确认
*/
type TaobaoWdkEquipmentConveyorStatusconfirmAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWdkEquipmentConveyorStatusconfirmResponse `json:"wdk_equipment_conveyor_statusconfirm_response,omitempty"` 
    TaobaoWdkEquipmentConveyorStatusconfirmResponse
}

/* model for simplify = false
type TaobaoWdkEquipmentConveyorStatusconfirmResponse struct {

    // result
    
    Result  *struct {
        TopWcsResult  *TopWcsResult `json:"top_wcs_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoWdkEquipmentConveyorStatusconfirmResponse struct {

    // result
    Result   *TopWcsResult `json:"result,omitempty"`

}
