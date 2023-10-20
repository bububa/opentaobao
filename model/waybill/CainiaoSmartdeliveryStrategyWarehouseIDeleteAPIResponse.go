package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse 删除智能发货引擎仓策略 API返回值
// cainiao.smartdelivery.strategy.warehouse.i.delete
//
// 删除智能发货引擎仓策略
type CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse struct {
	model.CommonResponse
	CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponseModel).Reset()
}

// CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponseModel is 删除智能发货引擎仓策略 成功返回结果
type CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_smartdelivery_strategy_warehouse_i_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	IsDeleteSuccess bool `json:"is_delete_success,omitempty" xml:"is_delete_success,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsDeleteSuccess = false
}

var poolCainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse)
	},
}

// GetCainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse 从 sync.Pool 获取 CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse
func GetCainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse() *CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse {
	return poolCainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse.Get().(*CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse)
}

// ReleaseCainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse 将 CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse 保存到 sync.Pool
func ReleaseCainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse(v *CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse) {
	v.Reset()
	poolCainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse.Put(v)
}
