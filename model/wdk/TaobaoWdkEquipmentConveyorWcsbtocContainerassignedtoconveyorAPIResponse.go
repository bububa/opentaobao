package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
容器被预分拣器分配到悬挂链 API返回值 
taobao.wdk.equipment.conveyor.wcsbtoc.containerassignedtoconveyor

容器被预分拣器分配到悬挂链
*/
type TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse struct {
    model.CommonResponse
    TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponseModel
}

// 容器被预分拣器分配到悬挂链 成功返回结果
type TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponseModel struct {
    XMLName xml.Name `xml:"wdk_equipment_conveyor_wcsbtoc_containerassignedtoconveyor_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TopWcsResult `json:"result,omitempty" xml:"result,omitempty"`
}
