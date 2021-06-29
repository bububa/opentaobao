package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
空间分组id查业务属性实例 API请求
alibaba.campus.space.group.getspacegroupwithattr

空间分组id查业务属性实例
*/
type AlibabaCampusSpaceGroupGetspacegroupwithattrRequest struct {
    model.Params
    // 操作用户上下文
    _context   *WorkBenchContext
    // 空间单元id
    _groupId   int64
}

// 初始化AlibabaCampusSpaceGroupGetspacegroupwithattrRequest对象
func NewAlibabaCampusSpaceGroupGetspacegroupwithattrRequest() *AlibabaCampusSpaceGroupGetspacegroupwithattrRequest{
    return &AlibabaCampusSpaceGroupGetspacegroupwithattrRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceGroupGetspacegroupwithattrRequest) GetApiMethodName() string {
    return "alibaba.campus.space.group.getspacegroupwithattr"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceGroupGetspacegroupwithattrRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Context Setter
// 操作用户上下文
func (r *AlibabaCampusSpaceGroupGetspacegroupwithattrRequest) SetContext(_context *WorkBenchContext) error {
    r._context = _context
    r.Set("context", _context)
    return nil
}

// Context Getter
func (r AlibabaCampusSpaceGroupGetspacegroupwithattrRequest) GetContext() *WorkBenchContext {
    return r._context
}
// GroupId Setter
// 空间单元id
func (r *AlibabaCampusSpaceGroupGetspacegroupwithattrRequest) SetGroupId(_groupId int64) error {
    r._groupId = _groupId
    r.Set("group_id", _groupId)
    return nil
}

// GroupId Getter
func (r AlibabaCampusSpaceGroupGetspacegroupwithattrRequest) GetGroupId() int64 {
    return r._groupId
}
