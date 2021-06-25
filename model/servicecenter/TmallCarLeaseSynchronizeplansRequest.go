package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步租赁方案 APIRequest
tmall.car.lease.synchronizeplans

租赁公司同步还款计划
*/
type TmallCarLeaseSynchronizeplansRequest struct {
    model.Params

    // 商品id
    itemId   int64 

    // 租赁计划
    plans   []CarLeasePlanDo 

}

func NewTmallCarLeaseSynchronizeplansRequest() *TmallCarLeaseSynchronizeplansRequest{
    return &TmallCarLeaseSynchronizeplansRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarLeaseSynchronizeplansRequest) GetApiMethodName() string {
    return "tmall.car.lease.synchronizeplans"
}

func (r TmallCarLeaseSynchronizeplansRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarLeaseSynchronizeplansRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TmallCarLeaseSynchronizeplansRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TmallCarLeaseSynchronizeplansRequest) SetPlans(plans []CarLeasePlanDo) error {
    r.plans = plans
    r.Set("plans", plans)
    return nil
}

func (r TmallCarLeaseSynchronizeplansRequest) GetPlans() []CarLeasePlanDo {
    return r.plans
}

