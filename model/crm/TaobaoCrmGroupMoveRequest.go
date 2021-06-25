package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分组移动 APIRequest
taobao.crm.group.move

将一个分组下的所有会员移动到另一个分组，会员从原分组中删除<br/>注：移动属性为异步任务建议先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。
*/
type TaobaoCrmGroupMoveRequest struct {
    model.Params

    // 需要移动的分组
    fromGroupId   int64 

    // 目的分组
    toGroupId   int64 

}

func NewTaobaoCrmGroupMoveRequest() *TaobaoCrmGroupMoveRequest{
    return &TaobaoCrmGroupMoveRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmGroupMoveRequest) GetApiMethodName() string {
    return "taobao.crm.group.move"
}

func (r TaobaoCrmGroupMoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCrmGroupMoveRequest) SetFromGroupId(fromGroupId int64) error {
    r.fromGroupId = fromGroupId
    r.Set("from_group_id", fromGroupId)
    return nil
}

func (r TaobaoCrmGroupMoveRequest) GetFromGroupId() int64 {
    return r.fromGroupId
}

func (r *TaobaoCrmGroupMoveRequest) SetToGroupId(toGroupId int64) error {
    r.toGroupId = toGroupId
    r.Set("to_group_id", toGroupId)
    return nil
}

func (r TaobaoCrmGroupMoveRequest) GetToGroupId() int64 {
    return r.toGroupId
}

