package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据分组ID查询相关的空间分组信息 APIRequest
alibaba.campus.space.group.getbyid

根据分组ID查询相关的空间分组信息
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
HSF方法名称：getById
*/
type AlibabaCampusSpaceGroupGetbyidRequest struct {
    model.Params

    // 用户环境
    param0   *WorkBenchContext 

    // 分组ID
    param1   int64 

}

func NewAlibabaCampusSpaceGroupGetbyidRequest() *AlibabaCampusSpaceGroupGetbyidRequest{
    return &AlibabaCampusSpaceGroupGetbyidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusSpaceGroupGetbyidRequest) GetApiMethodName() string {
    return "alibaba.campus.space.group.getbyid"
}

func (r AlibabaCampusSpaceGroupGetbyidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusSpaceGroupGetbyidRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaCampusSpaceGroupGetbyidRequest) GetParam0() *WorkBenchContext {
    return r.param0
}

func (r *AlibabaCampusSpaceGroupGetbyidRequest) SetParam1(param1 int64) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaCampusSpaceGroupGetbyidRequest) GetParam1() int64 {
    return r.param1
}

