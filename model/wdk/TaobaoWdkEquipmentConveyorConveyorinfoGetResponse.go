package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取五道口悬挂链信息 APIResponse
taobao.wdk.equipment.conveyor.conveyorinfo.get

获取五道口悬挂链信息
*/
type TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoWdkEquipmentConveyorConveyorinfoGetResponse
}

type TaobaoWdkEquipmentConveyorConveyorinfoGetResponse struct {
    XMLName xml.Name `xml:"wdk_equipment_conveyor_conveyorinfo_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值
    
    Result   *TaobaoWdkEquipmentConveyorConveyorinfoGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
