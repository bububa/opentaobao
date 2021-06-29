package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据分组ID查询相应的空间单元 APIRequest
alibaba.campus.space.unit.getlistbygroupid

根据分组ID查询相应的空间单元
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：getListByGroupId
*/
type AlibabaCampusSpaceUnitGetlistbygroupidRequest struct {
    model.Params

    // 分组ID
    param0   *WorkBenchContext 

    // 分组ID
    param1   int64 

}

func NewAlibabaCampusSpaceUnitGetlistbygroupidRequest() *AlibabaCampusSpaceUnitGetlistbygroupidRequest{
    return &AlibabaCampusSpaceUnitGetlistbygroupidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusSpaceUnitGetlistbygroupidRequest) GetApiMethodName() string {
    return "alibaba.campus.space.unit.getlistbygroupid"
}

func (r AlibabaCampusSpaceUnitGetlistbygroupidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusSpaceUnitGetlistbygroupidRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaCampusSpaceUnitGetlistbygroupidRequest) GetParam0() *WorkBenchContext {
    return r.param0
}

func (r *AlibabaCampusSpaceUnitGetlistbygroupidRequest) SetParam1(param1 int64) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaCampusSpaceUnitGetlistbygroupidRequest) GetParam1() int64 {
    return r.param1
}

