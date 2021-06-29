package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增业务属性实例接口 APIRequest
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

func NewAlibabaCampusSpaceAttrSetattrRequest() *AlibabaCampusSpaceAttrSetattrRequest{
    return &AlibabaCampusSpaceAttrSetattrRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusSpaceAttrSetattrRequest) GetApiMethodName() string {
    return "alibaba.campus.space.attr.setattr"
}

func (r AlibabaCampusSpaceAttrSetattrRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusSpaceAttrSetattrRequest) SetContext(context *WorkBenchContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

func (r AlibabaCampusSpaceAttrSetattrRequest) GetContext() *WorkBenchContext {
    return r.context
}

func (r *AlibabaCampusSpaceAttrSetattrRequest) SetList(list []TypeAttrInstanceRequest) error {
    r.list = list
    r.Set("list", list)
    return nil
}

func (r AlibabaCampusSpaceAttrSetattrRequest) GetList() []TypeAttrInstanceRequest {
    return r.list
}

