package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
异常通道日志查询 APIResponse
taobao.wdk.equipment.conveyor.exceptionslidewaylog.get

五道口悬挂链异常通道事件查询
*/
type TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResponse `json:"wdk_equipment_conveyor_exceptionslidewaylog_get_response,omitempty"` 
    TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResponse
}

/* model for simplify = false
type TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResponse struct {

    // 返回值
    
    Result  *struct {
        TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult  *TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult `json:"taobao_wdk_equipment_conveyor_exceptionslidewaylog_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResponse struct {

    // 返回值
    Result   *TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult `json:"result,omitempty"`

}
