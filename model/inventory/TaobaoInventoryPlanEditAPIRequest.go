package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoinventoryplaneditAPIRequest 设置计划库存 API请求
// taobao.inventory.plan.edit
//
// 初始化计划库存，或者编辑已经存在的计划库存
type TaobaoinventoryplaneditAPIRequest struct {
	model.Params
	// 计划库存设置入参
	_planTop *PlanTopDto
}

// NewTaobaoinventoryplaneditRequest 初始化TaobaoinventoryplaneditAPIRequest对象
func NewTaobaoinventoryplaneditRequest() *TaobaoinventoryplaneditAPIRequest {
	return &TaobaoinventoryplaneditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoinventoryplaneditAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.plan.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoinventoryplaneditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoinventoryplaneditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPlanTop is PlanTop Setter
// 计划库存设置入参
func (r *TaobaoinventoryplaneditAPIRequest) SetPlanTop(_planTop *PlanTopDto) error {
	r._planTop = _planTop
	r.Set("plan_top", _planTop)
	return nil
}

// GetPlanTop PlanTop Getter
func (r TaobaoinventoryplaneditAPIRequest) GetPlanTop() *PlanTopDto {
	return r._planTop
}
