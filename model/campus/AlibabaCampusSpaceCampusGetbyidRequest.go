package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据园区id获取园区信息 APIRequest
alibaba.campus.space.campus.getbyid

根据园区id获取园区信息
HSF接口名称：com.alibaba.campus.api.space.service.top.CampusApiTopService
HSF方法名称：getCampusById
*/
type AlibabaCampusSpaceCampusGetbyidRequest struct {
    model.Params

    // 园区ID
    param0   *WorkBenchContext 

    // 园区ID
    param1   int64 

}

func NewAlibabaCampusSpaceCampusGetbyidRequest() *AlibabaCampusSpaceCampusGetbyidRequest{
    return &AlibabaCampusSpaceCampusGetbyidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusSpaceCampusGetbyidRequest) GetApiMethodName() string {
    return "alibaba.campus.space.campus.getbyid"
}

func (r AlibabaCampusSpaceCampusGetbyidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusSpaceCampusGetbyidRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaCampusSpaceCampusGetbyidRequest) GetParam0() *WorkBenchContext {
    return r.param0
}

func (r *AlibabaCampusSpaceCampusGetbyidRequest) SetParam1(param1 int64) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaCampusSpaceCampusGetbyidRequest) GetParam1() int64 {
    return r.param1
}

