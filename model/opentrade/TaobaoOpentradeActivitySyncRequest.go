package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
尖货交易活动信息同步 APIRequest
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

func NewTaobaoOpentradeActivitySyncRequest() *TaobaoOpentradeActivitySyncRequest{
    return &TaobaoOpentradeActivitySyncRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpentradeActivitySyncRequest) GetApiMethodName() string {
    return "taobao.opentrade.activity.sync"
}

func (r TaobaoOpentradeActivitySyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpentradeActivitySyncRequest) SetActivityId(activityId string) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoOpentradeActivitySyncRequest) GetActivityId() string {
    return r.activityId
}

func (r *TaobaoOpentradeActivitySyncRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoOpentradeActivitySyncRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoOpentradeActivitySyncRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoOpentradeActivitySyncRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoOpentradeActivitySyncRequest) SetActivityName(activityName string) error {
    r.activityName = activityName
    r.Set("activity_name", activityName)
    return nil
}

func (r TaobaoOpentradeActivitySyncRequest) GetActivityName() string {
    return r.activityName
}

func (r *TaobaoOpentradeActivitySyncRequest) SetItemIdList(itemIdList []int64) error {
    r.itemIdList = itemIdList
    r.Set("item_id_list", itemIdList)
    return nil
}

func (r TaobaoOpentradeActivitySyncRequest) GetItemIdList() []int64 {
    return r.itemIdList
}

