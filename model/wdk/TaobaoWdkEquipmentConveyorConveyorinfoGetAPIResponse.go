package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse
获取五道口悬挂链信息 API返回值
taobao.wdk.equipment.conveyor.conveyorinfo.get

获取五道口悬挂链信息 */
type TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponseModel
}

// TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponseModel is 获取五道口悬挂链信息 成功返回结果
type TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_conveyor_conveyorinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TaobaoWdkEquipmentConveyorConveyorinfoGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
