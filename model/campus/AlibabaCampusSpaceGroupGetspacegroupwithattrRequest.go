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
    context   *WorkBenchContext
    // 空间单元id
    groupId   int64
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
func (r *AlibabaCampusSpaceGroupGetspacegroupwithattrRequest) SetContext(context *WorkBenchContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

// Context Getter
func (r AlibabaCampusSpaceGroupGetspacegroupwithattrRequest) GetContext() *WorkBenchContext {
    return r.context
}
// GroupId Setter
// 空间单元id
func (r *AlibabaCampusSpaceGroupGetspacegroupwithattrRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

// GroupId Getter
func (r AlibabaCampusSpaceGroupGetspacegroupwithattrRequest) GetGroupId() int64 {
    return r.groupId
}
