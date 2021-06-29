package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
空间分组id查业务属性实例 APIRequest
alibaba.campus.space.group.getspacegroupwithattr

空间分组id查业务属性实例
*/
type AlibabaCampusSpaceGroupGetspacegroupwithattrRequest struct {
    model.Params

    // 操作用户上下文
    context   *WorkBenchContext 

    // 空间单元id
    groupId   int64 

}

func NewAlibabaCampusSpaceGroupGetspacegroupwithattrRequest() *AlibabaCampusSpaceGroupGetspacegroupwithattrRequest{
    return &AlibabaCampusSpaceGroupGetspacegroupwithattrRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusSpaceGroupGetspacegroupwithattrRequest) GetApiMethodName() string {
    return "alibaba.campus.space.group.getspacegroupwithattr"
}

func (r AlibabaCampusSpaceGroupGetspacegroupwithattrRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusSpaceGroupGetspacegroupwithattrRequest) SetContext(context *WorkBenchContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

func (r AlibabaCampusSpaceGroupGetspacegroupwithattrRequest) GetContext() *WorkBenchContext {
    return r.context
}

func (r *AlibabaCampusSpaceGroupGetspacegroupwithattrRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r AlibabaCampusSpaceGroupGetspacegroupwithattrRequest) GetGroupId() int64 {
    return r.groupId
}

