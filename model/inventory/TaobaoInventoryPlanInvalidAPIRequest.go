package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoinventoryplaninvalidAPIRequest 失效计划库存 API请求
// taobao.inventory.plan.invalid
//
// 计划库存的失效服务
type TaobaoinventoryplaninvalidAPIRequest struct {
	model.Params
	// 计划库存失效入参
	_planStopTop *PlanStopTopDto
}

// NewTaobaoinventoryplaninvalidRequest 初始化TaobaoinventoryplaninvalidAPIRequest对象
func NewTaobaoinventoryplaninvalidRequest() *TaobaoinventoryplaninvalidAPIRequest {
	return &TaobaoinventoryplaninvalidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoinventoryplaninvalidAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.plan.invalid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoinventoryplaninvalidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoinventoryplaninvalidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPlanStopTop is PlanStopTop Setter
// 计划库存失效入参
func (r *TaobaoinventoryplaninvalidAPIRequest) SetPlanStopTop(_planStopTop *PlanStopTopDto) error {
	r._planStopTop = _planStopTop
	r.Set("plan_stop_top", _planStopTop)
	return nil
}

// GetPlanStopTop PlanStopTop Getter
func (r TaobaoinventoryplaninvalidAPIRequest) GetPlanStopTop() *PlanStopTopDto {
	return r._planStopTop
}
