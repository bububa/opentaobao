package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
给一批会员添加一个分组 APIRequest
taobao.crm.members.group.batchadd

为一批会员添加分组，接口返回添加是否成功,如至少有一个会员的分组添加成功，接口就返回成功，否则返回失败，如果当前会员已经拥有当前分组，则直接跳过
*/
type TaobaoCrmMembersGroupBatchaddRequest struct {
    model.Params

    // 分组id
    groupIds   []int64 

    // 买家昵称列表
    buyerNicks   []string 

}

func NewTaobaoCrmMembersGroupBatchaddRequest() *TaobaoCrmMembersGroupBatchaddRequest{
    return &TaobaoCrmMembersGroupBatchaddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmMembersGroupBatchaddRequest) GetApiMethodName() string {
    return "taobao.crm.members.group.batchadd"
}

func (r TaobaoCrmMembersGroupBatchaddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCrmMembersGroupBatchaddRequest) SetGroupIds(groupIds []int64) error {
    r.groupIds = groupIds
    r.Set("group_ids", groupIds)
    return nil
}

func (r TaobaoCrmMembersGroupBatchaddRequest) GetGroupIds() []int64 {
    return r.groupIds
}

func (r *TaobaoCrmMembersGroupBatchaddRequest) SetBuyerNicks(buyerNicks []string) error {
    r.buyerNicks = buyerNicks
    r.Set("buyer_nicks", buyerNicks)
    return nil
}

func (r TaobaoCrmMembersGroupBatchaddRequest) GetBuyerNicks() []string {
    return r.buyerNicks
}

