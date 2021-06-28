package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
硬件状态变化日志查询 APIResponse
taobao.wdk.equipment.conveyor.hardwarestatuslog.get

硬件状态变化日志查询
*/
type TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wdk_equipment_conveyor_hardwarestatuslog_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   *TaobaoWdkEquipmentConveyorHardwarestatuslogGetResult `json:"result,omitempty" xml:"