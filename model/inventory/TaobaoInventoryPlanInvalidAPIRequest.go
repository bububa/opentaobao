package inventory

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoInventoryPlanInvalidAPIRequest) Reset() {
	r._planStopTop = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInventoryPlanInvalidAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.plan.invalid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInventoryPlanInvalidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoInventoryPlanInvalidAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoInventoryPlanInvalidAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoInventoryPlanInvalidRequest()
	},
}

// GetTaobaoInventoryPlanInvalidRequest 从 sync.Pool 获取 TaobaoInventoryPlanInvalidAPIRequest
func GetTaobaoInventoryPlanInvalidAPIRequest() *TaobaoInventoryPlanInvalidAPIRequest {
	return poolTaobaoInventoryPlanInvalidAPIRequest.Get().(*TaobaoInventoryPlanInvalidAPIRequest)
}

// ReleaseTaobaoInventoryPlanInvalidAPIRequest 将 TaobaoInventoryPlanInvalidAPIRequest 放入 sync.Pool
func ReleaseTaobaoInventoryPlanInvalidAPIRequest(v *TaobaoInventoryPlanInvalidAPIRequest) {
	v.Reset()
	poolTaobaoInventoryPlanInvalidAPIRequest.Put(v)
}
