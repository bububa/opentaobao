package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取批次或波次中容器的信息 APIResponse
taobao.wdk.equipment.conveyor.containerinfo.get

获取批次或波次中容器的信息
*/
type TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWdkEquipmentConveyorContainerinfoGetResponse `json:"wdk_equipment_conveyor_containerinfo_get_response,omitempty"` 
    TaobaoWdkEquipmentConveyorContainerinfoGetResponse
}

/* model for simplify = false
type TaobaoWdkEquipmentConveyorContainerinfoGetResponse struct {

    // 返回值
    
    Result  *struct {
        TaobaoWdkEquipmentConveyorContainerinfoGetResult  *TaobaoWdkEquipmentConveyorContainerinfoGetResult `json:"taobao_wdk_equipment_conveyor_containerinfo_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoWdkEquipmentConveyorContainerinfoGetResponse struct {

    // 返回值
    Result   *TaobaoWdkEquipmentConveyorContainerinfoGetResult `json:"result,omitempty"`

}
