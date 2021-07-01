package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
悬挂链状态回传确认 API返回值 
taobao.wdk.equipment.conveyor.statusconfirm

悬挂链状态回传确认
*/
type TaobaoWdkEquipmentConveyorStatusconfirmAPIResponse struct {
    model.CommonResponse
    TaobaoWdkEquipmentConveyorStatusconfirmAPIResponseModel
}

// 悬挂链状态回传确认 成功返回结果
type TaobaoWdkEquipmentConveyorStatusconfirmAPIResponseModel struct {
    XMLName xml.Name `xml:"wdk_equipment_conveyor_statusconfirm_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TopWcsResult `json:"result,omitempty" xml:"result,omitempty"`
}
