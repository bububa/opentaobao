package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询空间分组业务属性 API请求
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

// 初始化AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest对象
func NewAlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest() *AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest{
    return &AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest) GetApiMethodName() string {
    return "alibaba.campus.space.group.getspacegrouplistwithattr"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Context Setter
// 操作用户上下文
func (r *AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest) SetContext(context *WorkBenchContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

// Context Getter
func (r AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest) GetContext() *WorkBenchContext {
    return r.context
}
// GroupQuery Setter
// 查询对象
func (r *AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest) SetGroupQuery(groupQuery *SpaceGroupQuery) error {
    r.groupQuery = groupQuery
    r.Set("group_query", groupQuery)
    return nil
}

// GroupQuery Getter
func (r AlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest) GetGroupQuery() *SpaceGroupQuery {
    return r.groupQuery
}
