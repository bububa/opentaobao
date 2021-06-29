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
    context   *WorkBenchContext
    // 业务属性实例集合
    list   []TypeAttrInstanceRequest
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
func (r *AlibabaCampusSpaceAttrSetattrRequest) SetContext(context *WorkBenchContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

// Context Getter
func (r AlibabaCampusSpaceAttrSetattrRequest) GetContext() *WorkBenchContext {
    return r.context
}
// List Setter
// 业务属性实例集合
func (r *AlibabaCampusSpaceAttrSetattrRequest) SetList(list []TypeAttrInstanceRequest) error {
    r.list = list
    r.Set("list", list)
    return nil
}

// List Getter
func (r AlibabaCampusSpaceAttrSetattrRequest) GetList() []TypeAttrInstanceRequest {
    return r.list
}
