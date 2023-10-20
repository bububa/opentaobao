package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWarehouseOperationUpdateAPIResponse 仓作业能力新建/更新 API返回值
// taobao.logistics.warehouse.operation.update
//
// 仓作业能力新建/更新
type TaobaoLogisticsWarehouseOperationUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsWarehouseOperationUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsWarehouseOperationUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsWarehouseOperationUpdateAPIResponseModel).Reset()
}

// TaobaoLogisticsWarehouseOperationUpdateAPIResponseModel is 仓作业能力新建/更新 成功返回结果
type TaobaoLogisticsWarehouseOperationUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_warehouse_operation_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 仓作业能力新建/更新出参
	WarehouseOperationUpdateResponse *WarehouseOperationUpdateResponse `json:"warehouse_operation_update_response,omitempty" xml:"warehouse_operation_update_response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsWarehouseOperationUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WarehouseOperationUpdateResponse = nil
}

var poolTaobaoLogisticsWarehouseOperationUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsWarehouseOperationUpdateAPIResponse)
	},
}

// GetTaobaoLogisticsWarehouseOperationUpdateAPIResponse 从 sync.Pool 获取 TaobaoLogisticsWarehouseOperationUpdateAPIResponse
func GetTaobaoLogisticsWarehouseOperationUpdateAPIResponse() *TaobaoLogisticsWarehouseOperationUpdateAPIResponse {
	return poolTaobaoLogisticsWarehouseOperationUpdateAPIResponse.Get().(*TaobaoLogisticsWarehouseOperationUpdateAPIResponse)
}

// ReleaseTaobaoLogisticsWarehouseOperationUpdateAPIResponse 将 TaobaoLogisticsWarehouseOperationUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsWarehouseOperationUpdateAPIResponse(v *TaobaoLogisticsWarehouseOperationUpdateAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsWarehouseOperationUpdateAPIResponse.Put(v)
}
