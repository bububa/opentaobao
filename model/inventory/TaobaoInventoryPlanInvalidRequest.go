package inventory

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
失效计划库存 APIRequest
taobao.inventory.plan.invalid

计划库存的失效服务
*/
type TaobaoInventoryPlanInvalidRequest struct {
    model.Params

    // 计划库存失效入参
    planStopTop   *PlanStopTopDto 

}

func NewTaobaoInventoryPlanInvalidRequest() *TaobaoInventoryPlanInvalidRequest{
    return &TaobaoInventoryPlanInvalidRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoInventoryPlanInvalidRequest) GetApiMethodName() string {
    return "taobao.inventory.plan.invalid"
}

func (r TaobaoInventoryPlanInvalidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoInventoryPlanInvalidRequest) SetPlanStopTop(planStopTop *PlanStopTopDto) error {
    r.planStopTop = planStopTop
    r.Set("plan_stop_top", planStopTop)
    return nil
}

func (r TaobaoInventoryPlanInvalidRequest) GetPlanStopTop() *PlanStopTopDto {
    return r.planStopTop
}

