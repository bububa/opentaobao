package inventory

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设置计划库存 API请求
taobao.inventory.plan.edit

初始化计划库存，或者编辑已经存在的计划库存
*/
type TaobaoInventoryPlanEditRequest struct {
    model.Params
    // 计划库存设置入参
    planTop   *PlanTopDto
}

// 初始化TaobaoInventoryPlanEditRequest对象
func NewTaobaoInventoryPlanEditRequest() *TaobaoInventoryPlanEditRequest{
    return &TaobaoInventoryPlanEditRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoInventoryPlanEditRequest) GetApiMethodName() string {
    return "taobao.inventory.plan.edit"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoInventoryPlanEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PlanTop Setter
// 计划库存设置入参
func (r *TaobaoInventoryPlanEditRequest) SetPlanTop(planTop *PlanTopDto) error {
    r.planTop = planTop
    r.Set("plan_top", planTop)
    return nil
}

// PlanTop Getter
func (r TaobaoInventoryPlanEditRequest) GetPlanTop() *PlanTopDto {
    return r.planTop
}
