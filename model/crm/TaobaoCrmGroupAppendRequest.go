package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
将一个分组添加到另外一个分组 APIRequest
taobao.crm.group.append

将某分组下的所有会员添加到另一个分组,注：1.该操作为异步任务，建议先调用taobao.crm.grouptask.check 确保涉及分组上没有任务；2.若分组下某会员分组数超最大限额，则该会员不会被添加到新分组，同时不影响其余会员添加分组，接口调用依然返回成功。
*/
type TaobaoCrmGroupAppendRequest struct {
    model.Params

    // 添加的来源分组
    fromGroupId   int64 

    // 添加的目标分组
    toGroupId   int64 

}

func NewTaobaoCrmGroupAppendRequest() *TaobaoCrmGroupAppendRequest{
    return &TaobaoCrmGroupAppendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmGroupAppendRequest) GetApiMethodName() string {
    return "taobao.crm.group.append"
}

func (r TaobaoCrmGroupAppendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCrmGroupAppendRequest) SetFromGroupId(fromGroupId int64) error {
    r.fromGroupId = fromGroupId
    r.Set("from_group_id", fromGroupId)
    return nil
}

func (r TaobaoCrmGroupAppendRequest) GetFromGroupId() int64 {
    return r.fromGroupId
}

func (r *TaobaoCrmGroupAppendRequest) SetToGroupId(toGroupId int64) error {
    r.toGroupId = toGroupId
    r.Set("to_group_id", toGroupId)
    return nil
}

func (r TaobaoCrmGroupAppendRequest) GetToGroupId() int64 {
    return r.toGroupId
}

