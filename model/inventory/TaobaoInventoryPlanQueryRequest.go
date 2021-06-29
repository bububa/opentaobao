package inventory

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
计划库存查询 APIRequest
taobao.inventory.plan.query

计划库存查询
*/
type TaobaoInventoryPlanQueryRequest struct {
    model.Params

    // 计划库存查询入参
    param   *InvUnifyPlanTopQuerys 

}

func NewTaobaoInventoryPlanQueryRequest() *TaobaoInventoryPlanQueryRequest{
    return &TaobaoInventoryPlanQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoInventoryPlanQueryRequest) GetApiMethodName() string {
    return "taobao.inventory.plan.query"
}

func (r TaobaoInventoryPlanQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoInventoryPlanQueryRequest) SetParam(param *InvUnifyPlanTopQuerys) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r TaobaoInventoryPlanQueryRequest) GetParam() *InvUnifyPlanTopQuerys {
    return r.param
}

