package inventory

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryPlanQuantityIncreaseAPIRequest 计划库存的增量编辑 API请求
// taobao.inventory.plan.quantity.increase
//
// 计划库存的增量编辑
type TaobaoInventoryPlanQuantityIncreaseAPIRequest struct {
	model.Params
	// 增量编辑计划库存入参
	_planInvAdjustTop *PlanInvAdjustTopDto
}

// NewTaobaoInventoryPlanQuantityIncreaseRequest 初始化TaobaoInventoryPlanQuantityIncreaseAPIRequest对象
func NewTaobaoInventoryPlanQuantityIncreaseRequest() *TaobaoInventoryPlanQuantityIncreaseAPIRequest {
	return &TaobaoInventoryPlanQuantityIncreaseAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoInventoryPlanQuantityIncreaseAPIRequest) Reset() {
	r._planInvAdjustTop = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInventoryPlanQuantityIncreaseAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.plan.quantity.increase"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInventoryPlanQuantityIncreaseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoInventoryPlanQuantityIncreaseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPlanInvAdjustTop is PlanInvAdjustTop Setter
// 增量编辑计划库存入参
func (r *TaobaoInventoryPlanQuantityIncreaseAPIRequest) SetPlanInvAdjustTop(_planInvAdjustTop *PlanInvAdjustTopDto) error {
	r._planInvAdjustTop = _planInvAdjustTop
	r.Set("plan_inv_adjust_top", _planInvAdjustTop)
	return nil
}

// GetPlanInvAdjustTop PlanInvAdjustTop Getter
func (r TaobaoInventoryPlanQuantityIncreaseAPIRequest) GetPlanInvAdjustTop() *PlanInvAdjustTopDto {
	return r._planInvAdjustTop
}

var poolTaobaoInventoryPlanQuantityIncreaseAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoInventoryPlanQuantityIncreaseRequest()
	},
}

// GetTaobaoInventoryPlanQuantityIncreaseRequest 从 sync.Pool 获取 TaobaoInventoryPlanQuantityIncreaseAPIRequest
func GetTaobaoInventoryPlanQuantityIncreaseAPIRequest() *TaobaoInventoryPlanQuantityIncreaseAPIRequest {
	return poolTaobaoInventoryPlanQuantityIncreaseAPIRequest.Get().(*TaobaoInventoryPlanQuantityIncreaseAPIRequest)
}

// ReleaseTaobaoInventoryPlanQuantityIncreaseAPIRequest 将 TaobaoInventoryPlanQuantityIncreaseAPIRequest 放入 sync.Pool
func ReleaseTaobaoInventoryPlanQuantityIncreaseAPIRequest(v *TaobaoInventoryPlanQuantityIncreaseAPIRequest) {
	v.Reset()
	poolTaobaoInventoryPlanQuantityIncreaseAPIRequest.Put(v)
}
