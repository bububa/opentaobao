package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取悬挂链系统事件 APIResponse
taobao.wdk.equipment.conveyor.systemevent.get

五道口悬挂链系统事件查询
*/
type TaobaoWdkEquipmentConveyorSystemeventGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWdkEquipmentConveyorSystemeventGetResponse `json:"wdk_equipment_conveyor_systemevent_get_response,omitempty"` 
    TaobaoWdkEquipmentConveyorSystemeventGetResponse
}

/* model for simplify = false
type TaobaoWdkEquipmentConveyorSystemeventGetResponse struct {

    // 返回值
    
    Result  *struct {
        TaobaoWdkEquipmentConveyorSystemeventGetResult  *TaobaoWdkEquipmentConveyorSystemeventGetResult `json:"taobao_wdk_equipment_conveyor_systemevent_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoWdkEquipmentConveyorSystemeventGetResponse struct {

    // 返回值
    Result   *TaobaoWdkEquipmentConveyorSystemeventGetResult `json:"result,omitempty"`

}
