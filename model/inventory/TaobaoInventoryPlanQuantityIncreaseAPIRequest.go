package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoinventoryplanquantityincreaseAPIRequest 计划库存的增量编辑 API请求
// taobao.inventory.plan.quantity.increase
//
// 计划库存的增量编辑
type TaobaoinventoryplanquantityincreaseAPIRequest struct {
	model.Params
	// 增量编辑计划库存入参
	_planInvAdjustTop *PlanInvAdjustTopDto
}

// NewTaobaoinventoryplanquantityincreaseRequest 初始化TaobaoinventoryplanquantityincreaseAPIRequest对象
func NewTaobaoinventoryplanquantityincreaseRequest() *TaobaoinventoryplanquantityincreaseAPIRequest {
	return &TaobaoinventoryplanquantityincreaseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoinventoryplanquantityincreaseAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.plan.quantity.increase"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoinventoryplanquantityincreaseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoinventoryplanquantityincreaseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPlanInvAdjustTop is PlanInvAdjustTop Setter
// 增量编辑计划库存入参
func (r *TaobaoinventoryplanquantityincreaseAPIRequest) SetPlanInvAdjustTop(_planInvAdjustTop *PlanInvAdjustTopDto) error {
	r._planInvAdjustTop = _planInvAdjustTop
	r.Set("plan_inv_adjust_top", _planInvAdjustTop)
	return nil
}

// GetPlanInvAdjustTop PlanInvAdjustTop Getter
func (r TaobaoinventoryplanquantityincreaseAPIRequest) GetPlanInvAdjustTop() *PlanInvAdjustTopDto {
	return r._planInvAdjustTop
}
