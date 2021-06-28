package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取五道口悬挂链信息 APIResponse
taobao.wdk.equipment.conveyor.conveyorinfo.get

获取五道口悬挂链信息
*/
type TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWdkEquipmentConveyorConveyorinfoGetResponse `json:"wdk_equipment_conveyor_conveyorinfo_get_response,omitempty"` 
    TaobaoWdkEquipmentConveyorConveyorinfoGetResponse
}

/* model for simplify = false
type TaobaoWdkEquipmentConveyorConveyorinfoGetResponse struct {

    // 返回值
    
    Result  *struct {
        TaobaoWdkEquipmentConveyorConveyorinfoGetResult  *TaobaoWdkEquipmentConveyorConveyorinfoGetResult `json:"taobao_wdk_equipment_conveyor_conveyorinfo_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoWdkEquipmentConveyorConveyorinfoGetResponse struct {

    // 返回值
    Result   *TaobaoWdkEquipmentConveyorConveyorinfoGetResult `json:"result,omitempty"`

}
