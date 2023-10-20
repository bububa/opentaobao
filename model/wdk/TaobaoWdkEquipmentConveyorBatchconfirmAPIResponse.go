package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorBatchconfirmAPIResponse 五道口悬挂链信息批量确认 API返回值
// taobao.wdk.equipment.conveyor.batchconfirm
//
// 批量消息确认
type TaobaoWdkEquipmentConveyorBatchconfirmAPIResponse struct {
	model.CommonResponse
	TaobaoWdkEquipmentConveyorBatchconfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentConveyorBatchconfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWdkEquipmentConveyorBatchconfirmAPIResponseModel).Reset()
}

// TaobaoWdkEquipmentConveyorBatchconfirmAPIResponseModel is 五道口悬挂链信息批量确认 成功返回结果
type TaobaoWdkEquipmentConveyorBatchconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_conveyor_batchconfirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopWcsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentConveyorBatchconfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWdkEquipmentConveyorBatchconfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWdkEquipmentConveyorBatchconfirmAPIResponse)
	},
}

// GetTaobaoWdkEquipmentConveyorBatchconfirmAPIResponse 从 sync.Pool 获取 TaobaoWdkEquipmentConveyorBatchconfirmAPIResponse
func GetTaobaoWdkEquipmentConveyorBatchconfirmAPIResponse() *TaobaoWdkEquipmentConveyorBatchconfirmAPIResponse {
	return poolTaobaoWdkEquipmentConveyorBatchconfirmAPIResponse.Get().(*TaobaoWdkEquipmentConveyorBatchconfirmAPIResponse)
}

// ReleaseTaobaoWdkEquipmentConveyorBatchconfirmAPIResponse 将 TaobaoWdkEquipmentConveyorBatchconfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWdkEquipmentConveyorBatchconfirmAPIResponse(v *TaobaoWdkEquipmentConveyorBatchconfirmAPIResponse) {
	v.Reset()
	poolTaobaoWdkEquipmentConveyorBatchconfirmAPIResponse.Put(v)
}
