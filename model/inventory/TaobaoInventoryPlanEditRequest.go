package inventory

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设置计划库存 APIRequest
taobao.inventory.plan.edit

初始化计划库存，或者编辑已经存在的计划库存
*/
type TaobaoInventoryPlanEditRequest struct {
    model.Params

    // 计划库存设置入参
    planTop   *PlanTopDto 

}

func NewTaobaoInventoryPlanEditRequest() *TaobaoInventoryPlanEditRequest{
    return &TaobaoInventoryPlanEditRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoInventoryPlanEditRequest) GetApiMethodName() string {
    return "taobao.inventory.plan.edit"
}

func (r TaobaoInventoryPlanEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoInventoryPlanEditRequest) SetPlanTop(planTop *PlanTopDto) error {
    r.planTop = planTop
    r.Set("plan_top", planTop)
    return nil
}

func (r TaobaoInventoryPlanEditRequest) GetPlanTop() *PlanTopDto {
    return r.planTop
}

