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
    Response *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorResponse `json:"taobao_wdk_equipment_conveyor_wcsbtoc_containerassignedtoconveyor_response,omitempty"`
}

type TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorResponse struct {

    // result
    Result   *TopWcsResult `json:"result,omitempty"`

}
