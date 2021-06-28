package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口悬挂链信息批量确认 APIResponse
taobao.wdk.equipment.conveyor.batchconfirm

批量消息确认
*/
type TaobaoWdkEquipmentConveyorBatchconfirmAPIResponse struct {
    model.CommonResponse
    TaobaoWdkEquipmentConveyorBatchconfirmResponse
}

type TaobaoWdkEquipmentConveyorBatchconfirmResponse struct {
    XMLName xml.Name `xml:"wdk_equipment_conveyor_batchconfirm_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TopWcsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
