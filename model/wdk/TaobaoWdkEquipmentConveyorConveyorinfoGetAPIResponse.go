package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse 获取五道口悬挂链信息 API返回值
// taobao.wdk.equipment.conveyor.conveyorinfo.get
//
// 获取五道口悬挂链信息
type TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponseModel).Reset()
}

// TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponseModel is 获取五道口悬挂链信息 成功返回结果
type TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_conveyor_conveyorinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TaobaoWdkEquipmentConveyorConveyorinfoGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse)
	},
}

// GetTaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse 从 sync.Pool 获取 TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse
func GetTaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse() *TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse {
	return poolTaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse.Get().(*TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse)
}

// ReleaseTaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse 将 TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse(v *TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse) {
	v.Reset()
	poolTaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse.Put(v)
}
