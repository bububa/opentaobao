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
type TmallCarLeaseSynchronizeplansAPIRequest struct {
    model.Params
    // 商品id
    _itemId   int64
    // 租赁计划
    _plans   []CarLeasePlanDO
}

// 初始化TmallCarLeaseSynchronizeplansAPIRequest对象
func NewTmallCarLeaseSynchronizeplansRequest() *TmallCarLeaseSynchronizeplansAPIRequest{
    return &TmallCarLeaseSynchronizeplansAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseSynchronizeplansAPIRequest) GetApiMethodName() string {
    return "tmall.car.lease.synchronizeplans"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseSynchronizeplansAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TmallCarLeaseSynchronizeplansAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TmallCarLeaseSynchronizeplansAPIRequest) GetItemId() int64 {
    return r._itemId
}
// Plans Setter
// 租赁计划
func (r *TmallCarLeaseSynchronizeplansAPIRequest) SetPlans(_plans []CarLeasePlanDO) error {
    r._plans = _plans
    r.Set("plans", _plans)
    return nil
}

// Plans Getter
func (r TmallCarLeaseSynchronizeplansAPIRequest) GetPlans() []CarLeasePlanDO {
    return r._plans
}
