package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIResponse
硬件状态变化日志查询 API返回值
taobao.wdk.equipment.conveyor.hardwarestatuslog.get

硬件状态变化日志查询 */
type TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIResponse struct {
	model.CommonResponse
	TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIResponseModel
}

// TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIResponseModel is 硬件状态变化日志查询 成功返回结果
type TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_conveyor_hardwarestatuslog_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TaobaoWdkEquipmentConveyorHardwarestatuslogGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
