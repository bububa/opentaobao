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
type AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest struct {
    model.Params
    // 操作用户上下文
    _context   *WorkBenchContext
    // 查询对象
    _groupQuery   *SpaceGroupQuery
}

// 初始化AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest对象
func NewAlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest() *AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest{
    return &AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.space.group.getspacegrouplistwithattr"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Context Setter
// 操作用户上下文
func (r *AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest) SetContext(_context *WorkBenchContext) error {
    r._context = _context
    r.Set("context", _context)
    return nil
}

// Context Getter
func (r AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest) GetContext() *WorkBenchContext {
    return r._context
}
// GroupQuery Setter
// 查询对象
func (r *AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest) SetGroupQuery(_groupQuery *SpaceGroupQuery) error {
    r._groupQuery = _groupQuery
    r.Set("group_query", _groupQuery)
    return nil
}

// GroupQuery Getter
func (r AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest) GetGroupQuery() *SpaceGroupQuery {
    return r._groupQuery
}
