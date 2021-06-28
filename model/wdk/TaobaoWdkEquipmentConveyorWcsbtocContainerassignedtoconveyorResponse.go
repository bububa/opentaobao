package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
容器被预分拣器分配到悬挂链 APIResponse
taobao.wdk.equipment.conveyor.wcsbtoc.containerassignedtoconveyor

容器被预分拣器分配到悬挂链
*/
type TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wdk_equipment_conveyor_wcsbtoc_containerassignedtoconveyor_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TopWcsResult `json:"result,omitempty" xml:"