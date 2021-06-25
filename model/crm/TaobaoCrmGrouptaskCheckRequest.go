package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询分组任务是否完成 APIRequest
taobao.crm.grouptask.check

检查一个分组上是否有异步任务,异步任务包括1.将一个分组下的所有用户添加到另外一个分组2.将一个分组下的所有用户移动到另外一个分组3.删除某个分组<br/>若分组上有任务则该属性不能被操作。
*/
type TaobaoCrmGrouptaskCheckRequest struct {
    model.Params

    // 分组id
    groupId   int64 

}

func NewTaobaoCrmGrouptaskCheckRequest() *TaobaoCrmGrouptaskCheckRequest{
    return &TaobaoCrmGrouptaskCheckRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmGrouptaskCheckRequest) GetApiMethodName() string {
    return "taobao.crm.grouptask.check"
}

func (r TaobaoCrmGrouptaskCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCrmGrouptaskCheckRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r TaobaoCrmGrouptaskCheckRequest) GetGroupId() int64 {
    return r.groupId
}

