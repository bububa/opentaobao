package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购获取用户参团信息 APIRequest
taobao.opentrade.group.member.info

组团购场景下，获取用户参团信息
*/
type TaobaoOpentradeGroupMemberInfoRequest struct {
    model.Params

    // 团id
    groupId   int64 

    // 用户openId
    openUserId   string 

}

func NewTaobaoOpentradeGroupMemberInfoRequest() *TaobaoOpentradeGroupMemberInfoRequest{
    return &TaobaoOpentradeGroupMemberInfoRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpentradeGroupMemberInfoRequest) GetApiMethodName() string {
    return "taobao.opentrade.group.member.info"
}

func (r TaobaoOpentradeGroupMemberInfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpentradeGroupMemberInfoRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r TaobaoOpentradeGroupMemberInfoRequest) GetGroupId() int64 {
    return r.groupId
}

func (r *TaobaoOpentradeGroupMemberInfoRequest) SetOpenUserId(openUserId string) error {
    r.openUserId = openUserId
    r.Set("open_user_id", openUserId)
    return nil
}

func (r TaobaoOpentradeGroupMemberInfoRequest) GetOpenUserId() string {
    return r.openUserId
}

