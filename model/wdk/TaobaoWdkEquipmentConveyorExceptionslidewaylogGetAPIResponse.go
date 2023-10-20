package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse 异常通道日志查询 API返回值
// taobao.wdk.equipment.conveyor.exceptionslidewaylog.get
//
// 五道口悬挂链异常通道事件查询
type TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse struct {
	model.CommonResponse
	TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponseModel).Reset()
}

// TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponseModel is 异常通道日志查询 成功返回结果
type TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_conveyor_exceptionslidewaylog_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse)
	},
}

// GetTaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse 从 sync.Pool 获取 TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse
func GetTaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse() *TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse {
	return poolTaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse.Get().(*TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse)
}

// ReleaseTaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse 将 TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse(v *TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse) {
	v.Reset()
	poolTaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse.Put(v)
}
