package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据ID查询指定空间单元信息 APIRequest
alibaba.campus.space.unit.getbyid

根据ID查询指定空间单元信息
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：getById
*/
type AlibabaCampusSpaceUnitGetbyidRequest struct {
    model.Params

    // 空间单元ID
    param0   *WorkBenchContext 

    // 空间单元ID
    param1   int64 

}

func NewAlibabaCampusSpaceUnitGetbyidRequest() *AlibabaCampusSpaceUnitGetbyidRequest{
    return &AlibabaCampusSpaceUnitGetbyidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusSpaceUnitGetbyidRequest) GetApiMethodName() string {
    return "alibaba.campus.space.unit.getbyid"
}

func (r AlibabaCampusSpaceUnitGetbyidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusSpaceUnitGetbyidRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaCampusSpaceUnitGetbyidRequest) GetParam0() *WorkBenchContext {
    return r.param0
}

func (r *AlibabaCampusSpaceUnitGetbyidRequest) SetParam1(param1 int64) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaCampusSpaceUnitGetbyidRequest) GetParam1() int64 {
    return r.param1
}

