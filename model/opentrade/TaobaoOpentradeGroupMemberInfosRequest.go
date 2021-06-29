package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购批量获取用户参团信息 API请求
taobao.opentrade.group.member.infos

组团购场景下，获取用户参团信息
*/
type TaobaoOpentradeGroupMemberInfosRequest struct {
    model.Params
    // 团id
    groupId   int64
    // 用户openId列表
    openUserIds   []string
}

// 初始化TaobaoOpentradeGroupMemberInfosRequest对象
func NewTaobaoOpentradeGroupMemberInfosRequest() *TaobaoOpentradeGroupMemberInfosRequest{
    return &TaobaoOpentradeGroupMemberInfosRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeGroupMemberInfosRequest) GetApiMethodName() string {
    return "taobao.opentrade.group.member.infos"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeGroupMemberInfosRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupId Setter
// 团id
func (r *TaobaoOpentradeGroupMemberInfosRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

// GroupId Getter
func (r TaobaoOpentradeGroupMemberInfosRequest) GetGroupId() int64 {
    return r.groupId
}
// OpenUserIds Setter
// 用户openId列表
func (r *TaobaoOpentradeGroupMemberInfosRequest) SetOpenUserIds(openUserIds []string) error {
    r.openUserIds = openUserIds
    r.Set("open_user_ids", openUserIds)
    return nil
}

// OpenUserIds Getter
func (r TaobaoOpentradeGroupMemberInfosRequest) GetOpenUserIds() []string {
    return r.openUserIds
}
