package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryPlanEditAPIRequest 设置计划库存 API请求
// taobao.inventory.plan.edit
//
// 初始化计划库存，或者编辑已经存在的计划库存
type TaobaoInventoryPlanEditAPIRequest struct {
	model.Params
	// 计划库存设置入参
	_planTop *PlanTopDto
}

// NewTaobaoInventoryPlanEditRequest 初始化TaobaoInventoryPlanEditAPIRequest对象
func NewTaobaoInventoryPlanEditRequest() *TaobaoInventoryPlanEditAPIRequest {
	return &TaobaoInventoryPlanEditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInventoryPlanEditAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.plan.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInventoryPlanEditAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PlanTop Setter
// 计划库存设置入参
func (r *TaobaoInventoryPlanEditAPIRequest) SetPlanTop(_planTop *PlanTopDto) error {
	r._planTop = _planTop
	r.Set("plan_top", _planTop)
	return nil
}

// Get PlanTop Getter
func (r TaobaoInventoryPlanEditAPIRequest) GetPlanTop() *PlanTopDto {
	return r._planTop
}
