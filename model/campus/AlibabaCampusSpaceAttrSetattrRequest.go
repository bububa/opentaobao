package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增业务属性实例接口 API请求
alibaba.campus.space.attr.setattr

新增业务属性实例接口
*/
type AlibabaCampusSpaceAttrSetattrRequest struct {
    model.Params
    // 操作用户上下文
    _context   *WorkBenchContext
    // 业务属性实例集合
    _list   []TypeAttrInstanceRequest
}

// 初始化AlibabaCampusSpaceAttrSetattrRequest对象
func NewAlibabaCampusSpaceAttrSetattrRequest() *AlibabaCampusSpaceAttrSetattrRequest{
    return &AlibabaCampusSpaceAttrSetattrRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceAttrSetattrRequest) GetApiMethodName() string {
    return "alibaba.campus.space.attr.setattr"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceAttrSetattrRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Context Setter
// 操作用户上下文
func (r *AlibabaCampusSpaceAttrSetattrRequest) SetContext(_context *WorkBenchContext) error {
    r._context = _context
    r.Set("context", _context)
    return nil
}

// Context Getter
func (r AlibabaCampusSpaceAttrSetattrRequest) GetContext() *WorkBenchContext {
    return r._context
}
// List Setter
// 业务属性实例集合
func (r *AlibabaCampusSpaceAttrSetattrRequest) SetList(_list []TypeAttrInstanceRequest) error {
    r._list = _list
    r.Set("list", _list)
    return nil
}

// List Getter
func (r AlibabaCampusSpaceAttrSetattrRequest) GetList() []TypeAttrInstanceRequest {
    return r._list
}
