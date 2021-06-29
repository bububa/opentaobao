package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购场景创建或更新组团活动 APIRequest
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

func NewTaobaoOpentradeGroupSyncRequest() *TaobaoOpentradeGroupSyncRequest{
    return &TaobaoOpentradeGroupSyncRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpentradeGroupSyncRequest) GetApiMethodName() string {
    return "taobao.opentrade.group.sync"
}

func (r TaobaoOpentradeGroupSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpentradeGroupSyncRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoOpentradeGroupSyncRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoOpentradeGroupSyncRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoOpentradeGroupSyncRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoOpentradeGroupSyncRequest) SetExpiration(expiration int64) error {
    r.expiration = expiration
    r.Set("expiration", expiration)
    return nil
}

func (r TaobaoOpentradeGroupSyncRequest) GetExpiration() int64 {
    return r.expiration
}

func (r *TaobaoOpentradeGroupSyncRequest) SetGoal(goal int64) error {
    r.goal = goal
    r.Set("goal", goal)
    return nil
}

func (r TaobaoOpentradeGroupSyncRequest) GetGoal() int64 {
    return r.goal
}

func (r *TaobaoOpentradeGroupSyncRequest) SetGroupType(groupType int64) error {
    r.groupType = groupType
    r.Set("group_type", groupType)
    return nil
}

func (r TaobaoOpentradeGroupSyncRequest) GetGroupType() int64 {
    return r.groupType
}

func (r *TaobaoOpentradeGroupSyncRequest) SetAllowType(allowType string) error {
    r.allowType = allowType
    r.Set("allow_type", allowType)
    return nil
}

func (r TaobaoOpentradeGroupSyncRequest) GetAllowType() string {
    return r.allowType
}

func (r *TaobaoOpentradeGroupSyncRequest) SetAllowWhiteList(allowWhiteList []string) error {
    r.allowWhiteList = allowWhiteList
    r.Set("allow_white_list", allowWhiteList)
    return nil
}

func (r TaobaoOpentradeGroupSyncRequest) GetAllowWhiteList() []string {
    return r.allowWhiteList
}

func (r *TaobaoOpentradeGroupSyncRequest) SetOpenLimit(openLimit int64) error {
    r.openLimit = openLimit
    r.Set("open_limit", openLimit)
    return nil
}

func (r TaobaoOpentradeGroupSyncRequest) GetOpenLimit() int64 {
    return r.openLimit
}

func (r *TaobaoOpentradeGroupSyncRequest) SetFailProcess(failProcess string) error {
    r.failProcess = failProcess
    r.Set("fail_process", failProcess)
    return nil
}

func (r TaobaoOpentradeGroupSyncRequest) GetFailProcess() string {
    return r.failProcess
}

func (r *TaobaoOpentradeGroupSyncRequest) SetDiscountPrice(discountPrice int64) error {
    r.discountPrice = discountPrice
    r.Set("discount_price", discountPrice)
    return nil
}

func (r TaobaoOpentradeGroupSyncRequest) GetDiscountPrice() int64 {
    return r.discountPrice
}

func (r *TaobaoOpentradeGroupSyncRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoOpentradeGroupSyncRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoOpentradeGroupSyncRequest) SetGroupActivityId(groupActivityId int64) error {
    r.groupActivityId = groupActivityId
    r.Set("group_activity_id", groupActivityId)
    return nil
}

func (r TaobaoOpentradeGroupSyncRequest) GetGroupActivityId() int64 {
    return r.groupActivityId
}

