package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增查询多个分组ID各自相关的空间单元信息 APIRequest
alibaba.campus.space.unit.getlistmapbygroupid

新增查询多个分组ID各自相关的空间单元信息
HSF接口名称：	com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：	getListMapByGroupIds
*/
type AlibabaCampusSpaceUnitGetlistmapbygroupidRequest struct {
    model.Params

    // 用户环境
    param0   *WorkBenchContext 

    // 查询封装
    param1   *SpaceUnitQuery 

}

func NewAlibabaCampusSpaceUnitGetlistmapbygroupidRequest() *AlibabaCampusSpaceUnitGetlistmapbygroupidRequest{
    return &AlibabaCampusSpaceUnitGetlistmapbygroupidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusSpaceUnitGetlistmapbygroupidRequest) GetApiMethodName() string {
    return "alibaba.campus.space.unit.getlistmapbygroupid"
}

func (r AlibabaCampusSpaceUnitGetlistmapbygroupidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusSpaceUnitGetlistmapbygroupidRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaCampusSpaceUnitGetlistmapbygroupidRequest) GetParam0() *WorkBenchContext {
    return r.param0
}

func (r *AlibabaCampusSpaceUnitGetlistmapbygroupidRequest) SetParam1(param1 *SpaceUnitQuery) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaCampusSpaceUnitGetlistmapbygroupidRequest) GetParam1() *SpaceUnitQuery {
    return r.param1
}

