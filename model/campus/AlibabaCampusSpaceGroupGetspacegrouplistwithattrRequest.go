package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询空间分组业务属性 APIRequest
alibaba.campus.space.group.getspacegrouplistwithattr

分页查询空间分组业务属性
*/
type AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest struct {
    model.Params

    // 操作用户上下文
    context   *WorkBenchContext 

    // 查询对象
    groupQuery   *SpaceGroupQuery 

}

func NewAlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest() *AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest{
    return &AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest) GetApiMethodName() string {
    return "alibaba.campus.space.group.getspacegrouplistwithattr"
}

func (r AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest) SetContext(context *WorkBenchContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

func (r AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest) GetContext() *WorkBenchContext {
    return r.context
}

func (r *AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest) SetGroupQuery(groupQuery *SpaceGroupQuery) error {
    r.groupQuery = groupQuery
    r.Set("group_query", groupQuery)
    return nil
}

func (r AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest) GetGroupQuery() *SpaceGroupQuery {
    return r.groupQuery
}

