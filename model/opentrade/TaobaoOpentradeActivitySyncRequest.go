package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
尖货交易活动信息同步 API请求
taobao.opentrade.activity.sync

尖货交易活动信息配置，创建或更新活动信息
在活动时间开始前，所有用户（包括标记可购买的用户），无法购买商品；
在活动时间内，标记可购买的用户可在小程序中跳转下单页，完成购买；
在活动结束后，对限购不再限制，平台开放购买，用户可在小程序内、商品详情、购物车下单购买；
*/
type TaobaoOpentradeActivitySyncRequest struct {
    model.Params
    // 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
    activityId   string
    // 活动开始时间
    startTime   string
    // 活动结束时间（全流程结束时间，非排队结束时间）
    endTime   string
    // 活动名称
    activityName   string
    // 活动关联的商品列表，使用逗号(,)分割
    itemIdList   []int64
}

// 初始化TaobaoOpentradeActivitySyncRequest对象
func NewTaobaoOpentradeActivitySyncRequest() *TaobaoOpentradeActivitySyncRequest{
    return &TaobaoOpentradeActivitySyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeActivitySyncRequest) GetApiMethodName() string {
    return "taobao.opentrade.activity.sync"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeActivitySyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
func (r *TaobaoOpentradeActivitySyncRequest) SetActivityId(activityId string) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoOpentradeActivitySyncRequest) GetActivityId() string {
    return r.activityId
}
// StartTime Setter
// 活动开始时间
func (r *TaobaoOpentradeActivitySyncRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoOpentradeActivitySyncRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 活动结束时间（全流程结束时间，非排队结束时间）
func (r *TaobaoOpentradeActivitySyncRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoOpentradeActivitySyncRequest) GetEndTime() string {
    return r.endTime
}
// ActivityName Setter
// 活动名称
func (r *TaobaoOpentradeActivitySyncRequest) SetActivityName(activityName string) error {
    r.activityName = activityName
    r.Set("activity_name", activityName)
    return nil
}

// ActivityName Getter
func (r TaobaoOpentradeActivitySyncRequest) GetActivityName() string {
    return r.activityName
}
// ItemIdList Setter
// 活动关联的商品列表，使用逗号(,)分割
func (r *TaobaoOpentradeActivitySyncRequest) SetItemIdList(itemIdList []int64) error {
    r.itemIdList = itemIdList
    r.Set("item_id_list", itemIdList)
    return nil
}

// ItemIdList Getter
func (r TaobaoOpentradeActivitySyncRequest) GetItemIdList() []int64 {
    return r.itemIdList
}
