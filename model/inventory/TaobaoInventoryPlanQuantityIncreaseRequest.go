package inventory

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
计划库存的增量编辑 APIRequest
taobao.inventory.plan.quantity.increase

计划库存的增量编辑
*/
type TaobaoInventoryPlanQuantityIncreaseRequest struct {
    model.Params

    // 增量编辑计划库存入参
    planInvAdjustTop   *PlanInvAdjustTopDto 

}

func NewTaobaoInventoryPlanQuantityIncreaseRequest() *TaobaoInventoryPlanQuantityIncreaseRequest{
    return &TaobaoInventoryPlanQuantityIncreaseRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoInventoryPlanQuantityIncreaseRequest) GetApiMethodName() string {
    return "taobao.inventory.plan.quantity.increase"
}

func (r TaobaoInventoryPlanQuantityIncreaseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoInventoryPlanQuantityIncreaseRequest) SetPlanInvAdjustTop(planInvAdjustTop *PlanInvAdjustTopDto) error {
    r.planInvAdjustTop = planInvAdjustTop
    r.Set("plan_inv_adjust_top", planInvAdjustTop)
    return nil
}

func (r TaobaoInventoryPlanQuantityIncreaseRequest) GetPlanInvAdjustTop() *PlanInvAdjustTopDto {
    return r.planInvAdjustTop
}

