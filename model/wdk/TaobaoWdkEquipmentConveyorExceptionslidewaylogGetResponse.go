package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
异常通道日志查询 APIResponse
taobao.wdk.equipment.conveyor.exceptionslidewaylog.get

五道口悬挂链异常通道事件查询
*/
type TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse struct {
    model.CommonResponse
    TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResponse
}

type TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResponse struct {
    XMLName xml.Name `xml:"wdk_equipment_conveyor_exceptionslidewaylog_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   *TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
