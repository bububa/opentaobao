package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse 智能发货引擎策略仓设置 API返回值
// cainiao.smartdelivery.strategy.warehouse.i.update
//
// 智能发货引擎发货策略设置仓维度
type CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse struct {
	model.CommonResponse
	CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponseModel).Reset()
}

// CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponseModel is 智能发货引擎策略仓设置 成功返回结果
type CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_smartdelivery_strategy_warehouse_i_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 仓信息
	WarehouseInfo *WarehouseDto `json:"warehouse_info,omitempty" xml:"warehouse_info,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WarehouseInfo = nil
}

var poolCainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse)
	},
}

// GetCainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse 从 sync.Pool 获取 CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse
func GetCainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse() *CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse {
	return poolCainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse.Get().(*CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse)
}

// ReleaseCainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse 将 CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse 保存到 sync.Pool
func ReleaseCainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse(v *CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse) {
	v.Reset()
	poolCainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse.Put(v)
}
