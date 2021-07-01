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
type TaobaoInventoryPlanQueryAPIRequest struct {
    model.Params
    // 计划库存查询入参
    _param   *InvUnifyPlanTopQuerys
}

// 初始化TaobaoInventoryPlanQueryAPIRequest对象
func NewTaobaoInventoryPlanQueryRequest() *TaobaoInventoryPlanQueryAPIRequest{
    return &TaobaoInventoryPlanQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoInventoryPlanQueryAPIRequest) GetApiMethodName() string {
    return "taobao.inventory.plan.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoInventoryPlanQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 计划库存查询入参
func (r *TaobaoInventoryPlanQueryAPIRequest) SetParam(_param *InvUnifyPlanTopQuerys) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r TaobaoInventoryPlanQueryAPIRequest) GetParam() *InvUnifyPlanTopQuerys {
    return r._param
}
