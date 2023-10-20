package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaowdkequipmentconveyorconveyorinfogetAPIResponse 获取五道口悬挂链信息 API返回值
// taobao.wdk.equipment.conveyor.conveyorinfo.get
//
// 获取五道口悬挂链信息
type TaobaowdkequipmentconveyorconveyorinfogetAPIResponse struct {
	model.CommonResponse
	TaobaowdkequipmentconveyorconveyorinfogetAPIResponseModel
}

// TaobaowdkequipmentconveyorconveyorinfogetAPIResponseModel is 获取五道口悬挂链信息 成功返回结果
type TaobaowdkequipmentconveyorconveyorinfogetAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_conveyor_conveyorinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TaobaowdkequipmentconveyorconveyorinfogetResult `json:"result,omitempty" xml:"result,omitempty"`
}
