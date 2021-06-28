package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
容器被预分拣器分配到悬挂链 APIResponse
taobao.wdk.equipment.conveyor.wcsbtoc.containerassignedtoconveyor

容器被预分拣器分配到悬挂链
*/
type TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorResponse `json:"wdk_equipment_conveyor_wcsbtoc_containerassignedtoconveyor_response,omitempty"` 
    TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorResponse
}

/* model for simplify = false
type TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorResponse struct {

    // result
    
    Result  *struct {
        TopWcsResult  *TopWcsResult `json:"top_wcs_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorResponse struct {

    // result
    Result   *TopWcsResult `json:"result,omitempty"`

}
