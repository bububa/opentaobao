package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购场景开团 APIRequest
taobao.opentrade.group.open

组团购场景下，团长开团
*/
type TaobaoOpentradeGroupOpenRequest struct {
    model.Params

    // 商品id
    itemId   int64 

    // 用户openId
    openUserId   string 

    // 组团活动id
    groupActivityId   int64 

}

func NewTaobaoOpentradeGroupOpenRequest() *TaobaoOpentradeGroupOpenRequest{
    return &TaobaoOpentradeGroupOpenRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpentradeGroupOpenRequest) GetApiMethodName() string {
    return "taobao.opentrade.group.open"
}

func (r TaobaoOpentradeGroupOpenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpentradeGroupOpenRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoOpentradeGroupOpenRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoOpentradeGroupOpenRequest) SetOpenUserId(openUserId string) error {
    r.openUserId = openUserId
    r.Set("open_user_id", openUserId)
    return nil
}

func (r TaobaoOpentradeGroupOpenRequest) GetOpenUserId() string {
    return r.openUserId
}

func (r *TaobaoOpentradeGroupOpenRequest) SetGroupActivityId(groupActivityId int64) error {
    r.groupActivityId = groupActivityId
    r.Set("group_activity_id", groupActivityId)
    return nil
}

func (r TaobaoOpentradeGroupOpenRequest) GetGroupActivityId() int64 {
    return r.groupActivityId
}

