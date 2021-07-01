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
type TaobaoInventoryPlanEditAPIRequest struct {
    model.Params
    // 计划库存设置入参
    _planTop   *PlanTopDTO
}

// 初始化TaobaoInventoryPlanEditAPIRequest对象
func NewTaobaoInventoryPlanEditRequest() *TaobaoInventoryPlanEditAPIRequest{
    return &TaobaoInventoryPlanEditAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoInventoryPlanEditAPIRequest) GetApiMethodName() string {
    return "taobao.inventory.plan.edit"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoInventoryPlanEditAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PlanTop Setter
// 计划库存设置入参
func (r *TaobaoInventoryPlanEditAPIRequest) SetPlanTop(_planTop *PlanTopDTO) error {
    r._planTop = _planTop
    r.Set("plan_top", _planTop)
    return nil
}

// PlanTop Getter
func (r TaobaoInventoryPlanEditAPIRequest) GetPlanTop() *PlanTopDTO {
    return r._planTop
}
