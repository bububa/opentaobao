package inventory

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
计划库存的增量编辑 API请求
taobao.inventory.plan.quantity.increase

计划库存的增量编辑
*/
type TaobaoInventoryPlanQuantityIncreaseRequest struct {
    model.Params
    // 增量编辑计划库存入参
    planInvAdjustTop   *PlanInvAdjustTopDto
}

// 初始化TaobaoInventoryPlanQuantityIncreaseRequest对象
func NewTaobaoInventoryPlanQuantityIncreaseRequest() *TaobaoInventoryPlanQuantityIncreaseRequest{
    return &TaobaoInventoryPlanQuantityIncreaseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoInventoryPlanQuantityIncreaseRequest) GetApiMethodName() string {
    return "taobao.inventory.plan.quantity.increase"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoInventoryPlanQuantityIncreaseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PlanInvAdjustTop Setter
// 增量编辑计划库存入参
func (r *TaobaoInventoryPlanQuantityIncreaseRequest) SetPlanInvAdjustTop(planInvAdjustTop *PlanInvAdjustTopDto) error {
    r.planInvAdjustTop = planInvAdjustTop
    r.Set("plan_inv_adjust_top", planInvAdjustTop)
    return nil
}

// PlanInvAdjustTop Getter
func (r TaobaoInventoryPlanQuantityIncreaseRequest) GetPlanInvAdjustTop() *PlanInvAdjustTopDto {
    return r.planInvAdjustTop
}
