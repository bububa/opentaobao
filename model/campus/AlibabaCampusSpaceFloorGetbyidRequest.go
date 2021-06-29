package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据id获取楼层 APIRequest
alibaba.campus.space.floor.getbyid

根据id获取楼层
*/
type AlibabaCampusSpaceFloorGetbyidRequest struct {
    model.Params

    // 环境上下文
    context   *WorkBenchContext 

    // 楼层id
    id   int64 

}

func NewAlibabaCampusSpaceFloorGetbyidRequest() *AlibabaCampusSpaceFloorGetbyidRequest{
    return &AlibabaCampusSpaceFloorGetbyidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusSpaceFloorGetbyidRequest) GetApiMethodName() string {
    return "alibaba.campus.space.floor.getbyid"
}

func (r AlibabaCampusSpaceFloorGetbyidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusSpaceFloorGetbyidRequest) SetContext(context *WorkBenchContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

func (r AlibabaCampusSpaceFloorGetbyidRequest) GetContext() *WorkBenchContext {
    return r.context
}

func (r *AlibabaCampusSpaceFloorGetbyidRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaCampusSpaceFloorGetbyidRequest) GetId() int64 {
    return r.id
}

