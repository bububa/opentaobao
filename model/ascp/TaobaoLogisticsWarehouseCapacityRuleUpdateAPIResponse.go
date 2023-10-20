package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse 仓产能创建/更新 API返回值
// taobao.logistics.warehouse.capacity.rule.update
//
// 仓产能创建/更新
type TaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponseModel).Reset()
}

// TaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponseModel is 仓产能创建/更新 成功返回结果
type TaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_warehouse_capacity_rule_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 仓产能创建/更新出参
	CapacityRuleUpdateResponse *CapacityRuleUpdateResponse `json:"capacity_rule_update_response,omitempty" xml:"capacity_rule_update_response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CapacityRuleUpdateResponse = nil
}

var poolTaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse)
	},
}

// GetTaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse 从 sync.Pool 获取 TaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse
func GetTaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse() *TaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse {
	return poolTaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse.Get().(*TaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse)
}

// ReleaseTaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse 将 TaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse(v *TaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse.Put(v)
}
