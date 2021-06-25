package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除分组 APIRequest
taobao.crm.group.delete

将该分组下的所有会员移除出该组，同时删除该分组。注：删除分组为异步任务，必须先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。
*/
type TaobaoCrmGroupDeleteRequest struct {
    model.Params

    // 要删除的分组id
    groupId   int64 

}

func NewTaobaoCrmGroupDeleteRequest() *TaobaoCrmGroupDeleteRequest{
    return &TaobaoCrmGroupDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmGroupDeleteRequest) GetApiMethodName() string {
    return "taobao.crm.group.delete"
}

func (r TaobaoCrmGroupDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCrmGroupDeleteRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r TaobaoCrmGroupDeleteRequest) GetGroupId() int64 {
    return r.groupId
}

