package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购获取用户参团信息 API请求
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

// 初始化TaobaoOpentradeGroupMemberInfoRequest对象
func NewTaobaoOpentradeGroupMemberInfoRequest() *TaobaoOpentradeGroupMemberInfoRequest{
    return &TaobaoOpentradeGroupMemberInfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeGroupMemberInfoRequest) GetApiMethodName() string {
    return "taobao.opentrade.group.member.info"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeGroupMemberInfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupId Setter
// 团id
func (r *TaobaoOpentradeGroupMemberInfoRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

// GroupId Getter
func (r TaobaoOpentradeGroupMemberInfoRequest) GetGroupId() int64 {
    return r.groupId
}
// OpenUserId Setter
// 用户openId
func (r *TaobaoOpentradeGroupMemberInfoRequest) SetOpenUserId(openUserId string) error {
    r.openUserId = openUserId
    r.Set("open_user_id", openUserId)
    return nil
}

// OpenUserId Getter
func (r TaobaoOpentradeGroupMemberInfoRequest) GetOpenUserId() string {
    return r.openUserId
}
