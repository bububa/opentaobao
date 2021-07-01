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
type TaobaoInventoryPlanQuantityIncreaseAPIRequest struct {
    model.Params
    // 增量编辑计划库存入参
    _planInvAdjustTop   *PlanInvAdjustTopDto
}

// 初始化TaobaoInventoryPlanQuantityIncreaseAPIRequest对象
func NewTaobaoInventoryPlanQuantityIncreaseRequest() *TaobaoInventoryPlanQuantityIncreaseAPIRequest{
    return &TaobaoInventoryPlanQuantityIncreaseAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoInventoryPlanQuantityIncreaseAPIRequest) GetApiMethodName() string {
    return "taobao.inventory.plan.quantity.increase"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoInventoryPlanQuantityIncreaseAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PlanInvAdjustTop Setter
// 增量编辑计划库存入参
func (r *TaobaoInventoryPlanQuantityIncreaseAPIRequest) SetPlanInvAdjustTop(_planInvAdjustTop *PlanInvAdjustTopDto) error {
    r._planInvAdjustTop = _planInvAdjustTop
    r.Set("plan_inv_adjust_top", _planInvAdjustTop)
    return nil
}

// PlanInvAdjustTop Getter
func (r TaobaoInventoryPlanQuantityIncreaseAPIRequest) GetPlanInvAdjustTop() *PlanInvAdjustTopDto {
    return r._planInvAdjustTop
}
