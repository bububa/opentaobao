package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryPlanInvalidAPIRequest 失效计划库存 API请求
// taobao.inventory.plan.invalid
//
// 计划库存的失效服务
type TaobaoInventoryPlanInvalidAPIRequest struct {
	model.Params
	// 计划库存失效入参
	_planStopTop *PlanStopTopDto
}

// NewTaobaoInventoryPlanInvalidRequest 初始化TaobaoInventoryPlanInvalidAPIRequest对象
func NewTaobaoInventoryPlanInvalidRequest() *TaobaoInventoryPlanInvalidAPIRequest {
	return &TaobaoInventoryPlanInvalidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInventoryPlanInvalidAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.plan.invalid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInventoryPlanInvalidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPlanStopTop is PlanStopTop Setter
// 计划库存失效入参
func (r *TaobaoInventoryPlanInvalidAPIRequest) SetPlanStopTop(_planStopTop *PlanStopTopDto) error {
	r._planStopTop = _planStopTop
	r.Set("plan_stop_top", _planStopTop)
	return nil
}

// GetPlanStopTop PlanStopTop Getter
func (r TaobaoInventoryPlanInvalidAPIRequest) GetPlanStopTop() *PlanStopTopDto {
	return r._planStopTop
}
