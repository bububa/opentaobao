package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步租赁方案 API请求
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

// 初始化TmallCarLeaseSynchronizeplansRequest对象
func NewTmallCarLeaseSynchronizeplansRequest() *TmallCarLeaseSynchronizeplansRequest{
    return &TmallCarLeaseSynchronizeplansRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseSynchronizeplansRequest) GetApiMethodName() string {
    return "tmall.car.lease.synchronizeplans"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseSynchronizeplansRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TmallCarLeaseSynchronizeplansRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TmallCarLeaseSynchronizeplansRequest) GetItemId() int64 {
    return r.itemId
}
// Plans Setter
// 租赁计划
func (r *TmallCarLeaseSynchronizeplansRequest) SetPlans(plans []CarLeasePlanDo) error {
    r.plans = plans
    r.Set("plans", plans)
    return nil
}

// Plans Getter
func (r TmallCarLeaseSynchronizeplansRequest) GetPlans() []CarLeasePlanDo {
    return r.plans
}
