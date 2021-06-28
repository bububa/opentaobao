package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取批次或波次中容器的信息 APIResponse
taobao.wdk.equipment.conveyor.containerinfo.get

获取批次或波次中容器的信息
*/
type TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wdk_equipment_conveyor_containerinfo_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   *TaobaoWdkEquipmentConveyorContainerinfoGetResult `json:"result,omitempty" xml:"