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
    _startTime   string
    // 组团活动结束时间
    _endTime   string
    // 成团有效期，单位为妙
    _expiration   int64
    // 成团的目标人数
    _goal   int64
    // 组团类型，0：拼团；1：团购
    _groupType   int64
    // 是否任何账号可开团。whitelist：仅白名单账号可开团  all：任何账号可开团
    _allowType   string
    // 允许开团的淘宝账号列表
    _allowWhiteList   []string
    // 组团类型为团购，可限制团长针对一个商品的开团数量上限
    _openLimit   int64
    // 未成团处理办法，close：系统关单；continue：订单继续下行
    _failProcess   string
    // 组团购买的折扣价，单位为分
    _discountPrice   int64
    // 商品ID
    _itemId   int64
    // 组团活动id
    _groupActivityId   int64
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
func (r *TaobaoOpentradeGroupSyncRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoOpentradeGroupSyncRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 组团活动结束时间
func (r *TaobaoOpentradeGroupSyncRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoOpentradeGroupSyncRequest) GetEndTime() string {
    return r._endTime
}
// Expiration Setter
// 成团有效期，单位为妙
func (r *TaobaoOpentradeGroupSyncRequest) SetExpiration(_expiration int64) error {
    r._expiration = _expiration
    r.Set("expiration", _expiration)
    return nil
}

// Expiration Getter
func (r TaobaoOpentradeGroupSyncRequest) GetExpiration() int64 {
    return r._expiration
}
// Goal Setter
// 成团的目标人数
func (r *TaobaoOpentradeGroupSyncRequest) SetGoal(_goal int64) error {
    r._goal = _goal
    r.Set("goal", _goal)
    return nil
}

// Goal Getter
func (r TaobaoOpentradeGroupSyncRequest) GetGoal() int64 {
    return r._goal
}
// GroupType Setter
// 组团类型，0：拼团；1：团购
func (r *TaobaoOpentradeGroupSyncRequest) SetGroupType(_groupType int64) error {
    r._groupType = _groupType
    r.Set("group_type", _groupType)
    return nil
}

// GroupType Getter
func (r TaobaoOpentradeGroupSyncRequest) GetGroupType() int64 {
    return r._groupType
}
// AllowType Setter
// 是否任何账号可开团。whitelist：仅白名单账号可开团  all：任何账号可开团
func (r *TaobaoOpentradeGroupSyncRequest) SetAllowType(_allowType string) error {
    r._allowType = _allowType
    r.Set("allow_type", _allowType)
    return nil
}

// AllowType Getter
func (r TaobaoOpentradeGroupSyncRequest) GetAllowType() string {
    return r._allowType
}
// AllowWhiteList Setter
// 允许开团的淘宝账号列表
func (r *TaobaoOpentradeGroupSyncRequest) SetAllowWhiteList(_allowWhiteList []string) error {
    r._allowWhiteList = _allowWhiteList
    r.Set("allow_white_list", _allowWhiteList)
    return nil
}

// AllowWhiteList Getter
func (r TaobaoOpentradeGroupSyncRequest) GetAllowWhiteList() []string {
    return r._allowWhiteList
}
// OpenLimit Setter
// 组团类型为团购，可限制团长针对一个商品的开团数量上限
func (r *TaobaoOpentradeGroupSyncRequest) SetOpenLimit(_openLimit int64) error {
    r._openLimit = _openLimit
    r.Set("open_limit", _openLimit)
    return nil
}

// OpenLimit Getter
func (r TaobaoOpentradeGroupSyncRequest) GetOpenLimit() int64 {
    return r._openLimit
}
// FailProcess Setter
// 未成团处理办法，close：系统关单；continue：订单继续下行
func (r *TaobaoOpentradeGroupSyncRequest) SetFailProcess(_failProcess string) error {
    r._failProcess = _failProcess
    r.Set("fail_process", _failProcess)
    return nil
}

// FailProcess Getter
func (r TaobaoOpentradeGroupSyncRequest) GetFailProcess() string {
    return r._failProcess
}
// DiscountPrice Setter
// 组团购买的折扣价，单位为分
func (r *TaobaoOpentradeGroupSyncRequest) SetDiscountPrice(_discountPrice int64) error {
    r._discountPrice = _discountPrice
    r.Set("discount_price", _discountPrice)
    return nil
}

// DiscountPrice Getter
func (r TaobaoOpentradeGroupSyncRequest) GetDiscountPrice() int64 {
    return r._discountPrice
}
// ItemId Setter
// 商品ID
func (r *TaobaoOpentradeGroupSyncRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpentradeGroupSyncRequest) GetItemId() int64 {
    return r._itemId
}
// GroupActivityId Setter
// 组团活动id
func (r *TaobaoOpentradeGroupSyncRequest) SetGroupActivityId(_groupActivityId int64) error {
    r._groupActivityId = _groupActivityId
    r.Set("group_activity_id", _groupActivityId)
    return nil
}

// GroupActivityId Getter
func (r TaobaoOpentradeGroupSyncRequest) GetGroupActivityId() int64 {
    return r._groupActivityId
}
