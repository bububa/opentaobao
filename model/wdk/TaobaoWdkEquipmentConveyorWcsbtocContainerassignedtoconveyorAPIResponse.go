package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse 容器被预分拣器分配到悬挂链 API返回值
// taobao.wdk.equipment.conveyor.wcsbtoc.containerassignedtoconveyor
//
// 容器被预分拣器分配到悬挂链
type TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse struct {
	model.CommonResponse
	TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponseModel).Reset()
}

// TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponseModel is 容器被预分拣器分配到悬挂链 成功返回结果
type TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_conveyor_wcsbtoc_containerassignedtoconveyor_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopWcsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse)
	},
}

// GetTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse 从 sync.Pool 获取 TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse
func GetTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse() *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse {
	return poolTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse.Get().(*TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse)
}

// ReleaseTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse 将 TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse(v *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse) {
	v.Reset()
	poolTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse.Put(v)
}
