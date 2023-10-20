package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaowdkequipmentconveyorcontainerinfogetAPIResponse 获取批次或波次中容器的信息 API返回值
// taobao.wdk.equipment.conveyor.containerinfo.get
//
// 获取批次或波次中容器的信息
type TaobaowdkequipmentconveyorcontainerinfogetAPIResponse struct {
	model.CommonResponse
	TaobaowdkequipmentconveyorcontainerinfogetAPIResponseModel
}

// TaobaowdkequipmentconveyorcontainerinfogetAPIResponseModel is 获取批次或波次中容器的信息 成功返回结果
type TaobaowdkequipmentconveyorcontainerinfogetAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_conveyor_containerinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TaobaowdkequipmentconveyorcontainerinfogetResult `json:"result,omitempty" xml:"result,omitempty"`
}
