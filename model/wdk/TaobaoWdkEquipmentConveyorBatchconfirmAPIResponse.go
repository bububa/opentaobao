package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口悬挂链信息批量确认 API返回值 
taobao.wdk.equipment.conveyor.batchconfirm

批量消息确认
*/
type TaobaoWdkEquipmentConveyorBatchconfirmAPIResponse struct {
    model.CommonResponse
    TaobaoWdkEquipmentConveyorBatchconfirmAPIResponseModel
}

// 五道口悬挂链信息批量确认 成功返回结果
type TaobaoWdkEquipmentConveyorBatchconfirmAPIResponseModel struct {
    XMLName xml.Name `xml:"wdk_equipment_conveyor_batchconfirm_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TopWcsResult `json:"result,omitempty" xml:"result,omitempty"`
}
