package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkEquipmentConveyorSystemeventGetAPIResponse
获取悬挂链系统事件 API返回值
taobao.wdk.equipment.conveyor.systemevent.get

五道口悬挂链系统事件查询 */
type TaobaoWdkEquipmentConveyorSystemeventGetAPIResponse struct {
	model.CommonResponse
	TaobaoWdkEquipmentConveyorSystemeventGetAPIResponseModel
}

// TaobaoWdkEquipmentConveyorSystemeventGetAPIResponseModel is 获取悬挂链系统事件 成功返回结果
type TaobaoWdkEquipmentConveyorSystemeventGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_conveyor_systemevent_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TaobaoWdkEquipmentConveyorSystemeventGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
