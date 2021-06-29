package inventory

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
计划库存查询 API请求
taobao.inventory.plan.query

计划库存查询
*/
type TaobaoInventoryPlanQueryRequest struct {
    model.Params
    // 计划库存查询入参
    param   *InvUnifyPlanTopQuerys
}

// 初始化TaobaoInventoryPlanQueryRequest对象
func NewTaobaoInventoryPlanQueryRequest() *TaobaoInventoryPlanQueryRequest{
    return &TaobaoInventoryPlanQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoInventoryPlanQueryRequest) GetApiMethodName() string {
    return "taobao.inventory.plan.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoInventoryPlanQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 计划库存查询入参
func (r *TaobaoInventoryPlanQueryRequest) SetParam(param *InvUnifyPlanTopQuerys) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r TaobaoInventoryPlanQueryRequest) GetParam() *InvUnifyPlanTopQuerys {
    return r.param
}
