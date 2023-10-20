package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse 获取批次或波次中容器的信息 API返回值
// taobao.wdk.equipment.conveyor.containerinfo.get
//
// 获取批次或波次中容器的信息
type TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponseModel).Reset()
}

// TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponseModel is 获取批次或波次中容器的信息 成功返回结果
type TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_conveyor_containerinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TaobaoWdkEquipmentConveyorContainerinfoGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse)
	},
}

// GetTaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse 从 sync.Pool 获取 TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse
func GetTaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse() *TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse {
	return poolTaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse.Get().(*TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse)
}

// ReleaseTaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse 将 TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse(v *TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse) {
	v.Reset()
	poolTaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse.Put(v)
}
