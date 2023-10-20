package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorStatusconfirmAPIResponse 悬挂链状态回传确认 API返回值
// taobao.wdk.equipment.conveyor.statusconfirm
//
// 悬挂链状态回传确认
type TaobaoWdkEquipmentConveyorStatusconfirmAPIResponse struct {
	model.CommonResponse
	TaobaoWdkEquipmentConveyorStatusconfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentConveyorStatusconfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWdkEquipmentConveyorStatusconfirmAPIResponseModel).Reset()
}

// TaobaoWdkEquipmentConveyorStatusconfirmAPIResponseModel is 悬挂链状态回传确认 成功返回结果
type TaobaoWdkEquipmentConveyorStatusconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_conveyor_statusconfirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopWcsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentConveyorStatusconfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWdkEquipmentConveyorStatusconfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWdkEquipmentConveyorStatusconfirmAPIResponse)
	},
}

// GetTaobaoWdkEquipmentConveyorStatusconfirmAPIResponse 从 sync.Pool 获取 TaobaoWdkEquipmentConveyorStatusconfirmAPIResponse
func GetTaobaoWdkEquipmentConveyorStatusconfirmAPIResponse() *TaobaoWdkEquipmentConveyorStatusconfirmAPIResponse {
	return poolTaobaoWdkEquipmentConveyorStatusconfirmAPIResponse.Get().(*TaobaoWdkEquipmentConveyorStatusconfirmAPIResponse)
}

// ReleaseTaobaoWdkEquipmentConveyorStatusconfirmAPIResponse 将 TaobaoWdkEquipmentConveyorStatusconfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWdkEquipmentConveyorStatusconfirmAPIResponse(v *TaobaoWdkEquipmentConveyorStatusconfirmAPIResponse) {
	v.Reset()
	poolTaobaoWdkEquipmentConveyorStatusconfirmAPIResponse.Put(v)
}
