package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取悬挂链系统事件 APIResponse
taobao.wdk.equipment.conveyor.systemevent.get

五道口悬挂链系统事件查询
*/
type TaobaoWdkEquipmentConveyorSystemeventGetAPIResponse struct {
    model.CommonResponse
    TaobaoWdkEquipmentConveyorSystemeventGetResponse
}

type TaobaoWdkEquipmentConveyorSystemeventGetResponse struct {
    XMLName xml.Name `xml:"wdk_equipment_conveyor_systemevent_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   *TaobaoWdkEquipmentConveyorSystemeventGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
