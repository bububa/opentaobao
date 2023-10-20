package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWarehouseCapacityRuleUpdateAPIRequest 仓产能创建/更新 API请求
// taobao.logistics.warehouse.capacity.rule.update
//
// 仓产能创建/更新
type TaobaoLogisticsWarehouseCapacityRuleUpdateAPIRequest struct {
	model.Params
	// 仓产能创建/更新入参
	_capacityRuleUpdateRequest *CapacityRuleUpdateRequest
}

// NewTaobaoLogisticsWarehouseCapacityRuleUpdateRequest 初始化TaobaoLogisticsWarehouseCapacityRuleUpdateAPIRequest对象
func NewTaobaoLogisticsWarehouseCapacityRuleUpdateRequest() *TaobaoLogisticsWarehouseCapacityRuleUpdateAPIRequest {
	return &TaobaoLogisticsWarehouseCapacityRuleUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsWarehouseCapacityRuleUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.warehouse.capacity.rule.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsWarehouseCapacityRuleUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsWarehouseCapacityRuleUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCapacityRuleUpdateRequest is CapacityRuleUpdateRequest Setter
// 仓产能创建/更新入参
func (r *TaobaoLogisticsWarehouseCapacityRuleUpdateAPIRequest) SetCapacityRuleUpdateRequest(_capacityRuleUpdateRequest *CapacityRuleUpdateRequest) error {
	r._capacityRuleUpdateRequest = _capacityRuleUpdateRequest
	r.Set("capacity_rule_update_request", _capacityRuleUpdateRequest)
	return nil
}

// GetCapacityRuleUpdateRequest CapacityRuleUpdateRequest Getter
func (r TaobaoLogisticsWarehouseCapacityRuleUpdateAPIRequest) GetCapacityRuleUpdateRequest() *CapacityRuleUpdateRequest {
	return r._capacityRuleUpdateRequest
}
