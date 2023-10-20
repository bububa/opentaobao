package inventory

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoInventoryPlanEditAPIRequest) Reset() {
	r._planTop = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInventoryPlanEditAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.plan.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInventoryPlanEditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoInventoryPlanEditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPlanTop is PlanTop Setter
// 计划库存设置入参
func (r *TaobaoInventoryPlanEditAPIRequest) SetPlanTop(_planTop *PlanTopDto) error {
	r._planTop = _planTop
	r.Set("plan_top", _planTop)
	return nil
}

// GetPlanTop PlanTop Getter
func (r TaobaoInventoryPlanEditAPIRequest) GetPlanTop() *PlanTopDto {
	return r._planTop
}

var poolTaobaoInventoryPlanEditAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoInventoryPlanEditRequest()
	},
}

// GetTaobaoInventoryPlanEditRequest 从 sync.Pool 获取 TaobaoInventoryPlanEditAPIRequest
func GetTaobaoInventoryPlanEditAPIRequest() *TaobaoInventoryPlanEditAPIRequest {
	return poolTaobaoInventoryPlanEditAPIRequest.Get().(*TaobaoInventoryPlanEditAPIRequest)
}

// ReleaseTaobaoInventoryPlanEditAPIRequest 将 TaobaoInventoryPlanEditAPIRequest 放入 sync.Pool
func ReleaseTaobaoInventoryPlanEditAPIRequest(v *TaobaoInventoryPlanEditAPIRequest) {
	v.Reset()
	poolTaobaoInventoryPlanEditAPIRequest.Put(v)
}
