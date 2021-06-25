package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
悬挂链状态回传确认 APIResponse
taobao.wdk.equipment.conveyor.statusconfirm

悬挂链状态回传确认
*/
type TaobaoWdkEquipmentConveyorStatusconfirmAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWdkEquipmentConveyorStatusconfirmResponse `json:"taobao_wdk_equipment_conveyor_statusconfirm_response,omitempty"`
}

type TaobaoWdkEquipmentConveyorStatusconfirmResponse struct {

    // result
    Result   *TopWcsResult `json:"result,omitempty"`

}
