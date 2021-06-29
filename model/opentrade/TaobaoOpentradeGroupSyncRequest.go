package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购场景创建或更新组团活动 API请求
taobao.opentrade.group.sync

组团购场景中创建团购活动
*/
type TaobaoOpentradeGroupSyncRequest struct {
    model.Params
    // 组团活动开始时间
    startTime   string
    // 组团活动结束时间
    endTime   string
    // 成团有效期，单位为妙
    expiration   int64
    // 成团的目标人数
    goal   int64
    // 组团类型，0：拼团；1：团购
    groupType   int64
    // 是否任何账号可开团。whitelist：仅白名单账号可开团  all：任何账号可开团
    allowType   string
    // 允许开团的淘宝账号列表
    allowWhiteList   []string
    // 组团类型为团购，可限制团长针对一个商品的开团数量上限
    openLimit   int64
    // 未成团处理办法，close：系统关单；continue：订单继续下行
    failProcess   string
    // 组团购买的折扣价，单位为分
    discountPrice   int64
    // 商品ID
    itemId   int64
    // 组团活动id
    groupActivityId   int64
}

// 初始化TaobaoOpentradeGroupSyncRequest对象
func NewTaobaoOpentradeGroupSyncRequest() *TaobaoOpentradeGroupSyncRequest{
    return &TaobaoOpentradeGroupSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeGroupSyncRequest) GetApiMethodName() string {
    return "taobao.opentrade.group.sync"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeGroupSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartTime Setter
// 组团活动开始时间
func (r *TaobaoOpentradeGroupSyncRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoOpentradeGroupSyncRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 组团活动结束时间
func (r *TaobaoOpentradeGroupSyncRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoOpentradeGroupSyncRequest) GetEndTime() string {
    return r.endTime
}
// Expiration Setter
// 成团有效期，单位为妙
func (r *TaobaoOpentradeGroupSyncRequest) SetExpiration(expiration int64) error {
    r.expiration = expiration
    r.Set("expiration", expiration)
    return nil
}

// Expiration Getter
func (r TaobaoOpentradeGroupSyncRequest) GetExpiration() int64 {
    return r.expiration
}
// Goal Setter
// 成团的目标人数
func (r *TaobaoOpentradeGroupSyncRequest) SetGoal(goal int64) error {
    r.goal = goal
    r.Set("goal", goal)
    return nil
}

// Goal Getter
func (r TaobaoOpentradeGroupSyncRequest) GetGoal() int64 {
    return r.goal
}
// GroupType Setter
// 组团类型，0：拼团；1：团购
func (r *TaobaoOpentradeGroupSyncRequest) SetGroupType(groupType int64) error {
    r.groupType = groupType
    r.Set("group_type", groupType)
    return nil
}

// GroupType Getter
func (r TaobaoOpentradeGroupSyncRequest) GetGroupType() int64 {
    return r.groupType
}
// AllowType Setter
// 是否任何账号可开团。whitelist：仅白名单账号可开团  all：任何账号可开团
func (r *TaobaoOpentradeGroupSyncRequest) SetAllowType(allowType string) error {
    r.allowType = allowType
    r.Set("allow_type", allowType)
    return nil
}

// AllowType Getter
func (r TaobaoOpentradeGroupSyncRequest) GetAllowType() string {
    return r.allowType
}
// AllowWhiteList Setter
// 允许开团的淘宝账号列表
func (r *TaobaoOpentradeGroupSyncRequest) SetAllowWhiteList(allowWhiteList []string) error {
    r.allowWhiteList = allowWhiteList
    r.Set("allow_white_list", allowWhiteList)
    return nil
}

// AllowWhiteList Getter
func (r TaobaoOpentradeGroupSyncRequest) GetAllowWhiteList() []string {
    return r.allowWhiteList
}
// OpenLimit Setter
// 组团类型为团购，可限制团长针对一个商品的开团数量上限
func (r *TaobaoOpentradeGroupSyncRequest) SetOpenLimit(openLimit int64) error {
    r.openLimit = openLimit
    r.Set("open_limit", openLimit)
    return nil
}

// OpenLimit Getter
func (r TaobaoOpentradeGroupSyncRequest) GetOpenLimit() int64 {
    return r.openLimit
}
// FailProcess Setter
// 未成团处理办法，close：系统关单；continue：订单继续下行
func (r *TaobaoOpentradeGroupSyncRequest) SetFailProcess(failProcess string) error {
    r.failProcess = failProcess
    r.Set("fail_process", failProcess)
    return nil
}

// FailProcess Getter
func (r TaobaoOpentradeGroupSyncRequest) GetFailProcess() string {
    return r.failProcess
}
// DiscountPrice Setter
// 组团购买的折扣价，单位为分
func (r *TaobaoOpentradeGroupSyncRequest) SetDiscountPrice(discountPrice int64) error {
    r.discountPrice = discountPrice
    r.Set("discount_price", discountPrice)
    return nil
}

// DiscountPrice Getter
func (r TaobaoOpentradeGroupSyncRequest) GetDiscountPrice() int64 {
    return r.discountPrice
}
// ItemId Setter
// 商品ID
func (r *TaobaoOpentradeGroupSyncRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpentradeGroupSyncRequest) GetItemId() int64 {
    return r.itemId
}
// GroupActivityId Setter
// 组团活动id
func (r *TaobaoOpentradeGroupSyncRequest) SetGroupActivityId(groupActivityId int64) error {
    r.groupActivityId = groupActivityId
    r.Set("group_activity_id", groupActivityId)
    return nil
}

// GroupActivityId Getter
func (r TaobaoOpentradeGroupSyncRequest) GetGroupActivityId() int64 {
    return r.groupActivityId
}
