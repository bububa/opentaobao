package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
硬件状态变化日志查询 APIResponse
taobao.wdk.equipment.conveyor.hardwarestatuslog.get

硬件状态变化日志查询
*/
type TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWdkEquipmentConveyorHardwarestatuslogGetResponse `json:"taobao_wdk_equipment_conveyor_hardwarestatuslog_get_response,omitempty"`
}

type TaobaoWdkEquipmentConveyorHardwarestatuslogGetResponse struct {

    // 返回值
    Result   *TaobaoWdkEquipmentConveyorHardwarestatuslogGetResult `json:"result,omitempty"`

}