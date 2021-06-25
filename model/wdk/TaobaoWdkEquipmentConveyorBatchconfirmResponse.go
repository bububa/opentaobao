package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
五道口悬挂链信息批量确认 APIResponse
taobao.wdk.equipment.conveyor.batchconfirm

批量消息确认
*/
type TaobaoWdkEquipmentConveyorBatchconfirmAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWdkEquipmentConveyorBatchconfirmResponse `json:"taobao_wdk_equipment_conveyor_batchconfirm_response,omitempty"`
}

type TaobaoWdkEquipmentConveyorBatchconfirmResponse struct {

    // result
    Result   *TopWcsResult `json:"result,omitempty"`

}
