package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaowdkequipmentconveyorstatusconfirmAPIResponse 悬挂链状态回传确认 API返回值
// taobao.wdk.equipment.conveyor.statusconfirm
//
// 悬挂链状态回传确认
type TaobaowdkequipmentconveyorstatusconfirmAPIResponse struct {
	model.CommonResponse
	TaobaowdkequipmentconveyorstatusconfirmAPIResponseModel
}

// TaobaowdkequipmentconveyorstatusconfirmAPIResponseModel is 悬挂链状态回传确认 成功返回结果
type TaobaowdkequipmentconveyorstatusconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_conveyor_statusconfirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopWcsResult `json:"result,omitempty" xml:"result,omitempty"`
}
