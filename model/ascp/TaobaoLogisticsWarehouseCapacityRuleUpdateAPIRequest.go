package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswarehousecapacityruleupdateAPIRequest 仓产能创建/更新 API请求
// taobao.logistics.warehouse.capacity.rule.update
//
// 仓产能创建/更新
type TaobaologisticswarehousecapacityruleupdateAPIRequest struct {
	model.Params
	// 仓产能创建/更新入参
	_capacityRuleUpdateRequest *CapacityRuleUpdateRequest
}

// NewTaobaologisticswarehousecapacityruleupdateRequest 初始化TaobaologisticswarehousecapacityruleupdateAPIRequest对象
func NewTaobaologisticswarehousecapacityruleupdateRequest() *TaobaologisticswarehousecapacityruleupdateAPIRequest {
	return &TaobaologisticswarehousecapacityruleupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticswarehousecapacityruleupdateAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.warehouse.capacity.rule.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticswarehousecapacityruleupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticswarehousecapacityruleupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCapacityRuleUpdateRequest is CapacityRuleUpdateRequest Setter
// 仓产能创建/更新入参
func (r *TaobaologisticswarehousecapacityruleupdateAPIRequest) SetCapacityRuleUpdateRequest(_capacityRuleUpdateRequest *CapacityRuleUpdateRequest) error {
	r._capacityRuleUpdateRequest = _capacityRuleUpdateRequest
	r.Set("capacity_rule_update_request", _capacityRuleUpdateRequest)
	return nil
}

// GetCapacityRuleUpdateRequest CapacityRuleUpdateRequest Getter
func (r TaobaologisticswarehousecapacityruleupdateAPIRequest) GetCapacityRuleUpdateRequest() *CapacityRuleUpdateRequest {
	return r._capacityRuleUpdateRequest
}
