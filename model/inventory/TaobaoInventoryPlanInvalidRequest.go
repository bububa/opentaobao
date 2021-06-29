package inventory

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
失效计划库存 API请求
taobao.inventory.plan.invalid

计划库存的失效服务
*/
type TaobaoInventoryPlanInvalidRequest struct {
    model.Params
    // 计划库存失效入参
    _planStopTop   *PlanStopTopDTO
}

// 初始化TaobaoInventoryPlanInvalidRequest对象
func NewTaobaoInventoryPlanInvalidRequest() *TaobaoInventoryPlanInvalidRequest{
    return &TaobaoInventoryPlanInvalidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoInventoryPlanInvalidRequest) GetApiMethodName() string {
    return "taobao.inventory.plan.invalid"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoInventoryPlanInvalidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PlanStopTop Setter
// 计划库存失效入参
func (r *TaobaoInventoryPlanInvalidRequest) SetPlanStopTop(_planStopTop *PlanStopTopDTO) error {
    r._planStopTop = _planStopTop
    r.Set("plan_stop_top", _planStopTop)
    return nil
}

// PlanStopTop Getter
func (r TaobaoInventoryPlanInvalidRequest) GetPlanStopTop() *PlanStopTopDTO {
    return r._planStopTop
}
