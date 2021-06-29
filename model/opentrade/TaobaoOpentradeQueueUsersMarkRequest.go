package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
尖货交易可购买用户标记 API请求
taobao.opentrade.queue.users.mark

尖货交易用户标记信息回传，根据openId标记用户可购买商品
*/
type TaobaoOpentradeQueueUsersMarkRequest struct {
    model.Params
    // 用户状态，可任意传入，后续查询返回
    status   string
    // 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
    activityId   string
    // 排队商品SKU ID，不存在传0
    skuId   int64
    // 排队商品ID
    itemId   int64
    // 本次待标记的用户列表，多个以逗号(,)分割，最大20个
    openUserIds   []string
    // 是否目标用户，传入true后，用户可购买商品
    hit   bool
}

// 初始化TaobaoOpentradeQueueUsersMarkRequest对象
func NewTaobaoOpentradeQueueUsersMarkRequest() *TaobaoOpentradeQueueUsersMarkRequest{
    return &TaobaoOpentradeQueueUsersMarkRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeQueueUsersMarkRequest) GetApiMethodName() string {
    return "taobao.opentrade.queue.users.mark"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeQueueUsersMarkRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// 用户状态，可任意传入，后续查询返回
func (r *TaobaoOpentradeQueueUsersMarkRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoOpentradeQueueUsersMarkRequest) GetStatus() string {
    return r.status
}
// ActivityId Setter
// 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
func (r *TaobaoOpentradeQueueUsersMarkRequest) SetActivityId(activityId string) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoOpentradeQueueUsersMarkRequest) GetActivityId() string {
    return r.activityId
}
// SkuId Setter
// 排队商品SKU ID，不存在传0
func (r *TaobaoOpentradeQueueUsersMarkRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

// SkuId Getter
func (r TaobaoOpentradeQueueUsersMarkRequest) GetSkuId() int64 {
    return r.skuId
}
// ItemId Setter
// 排队商品ID
func (r *TaobaoOpentradeQueueUsersMarkRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpentradeQueueUsersMarkRequest) GetItemId() int64 {
    return r.itemId
}
// OpenUserIds Setter
// 本次待标记的用户列表，多个以逗号(,)分割，最大20个
func (r *TaobaoOpentradeQueueUsersMarkRequest) SetOpenUserIds(openUserIds []string) error {
    r.openUserIds = openUserIds
    r.Set("open_user_ids", openUserIds)
    return nil
}

// OpenUserIds Getter
func (r TaobaoOpentradeQueueUsersMarkRequest) GetOpenUserIds() []string {
    return r.openUserIds
}
// Hit Setter
// 是否目标用户，传入true后，用户可购买商品
func (r *TaobaoOpentradeQueueUsersMarkRequest) SetHit(hit bool) error {
    r.hit = hit
    r.Set("hit", hit)
    return nil
}

// Hit Getter
func (r TaobaoOpentradeQueueUsersMarkRequest) GetHit() bool {
    return r.hit
}
