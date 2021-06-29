package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据园区id及TypeId获取空间分组 APIRequest
alibaba.campus.space.group.getlistbycampusandtype

根据园区id及TypeId获取空间分组
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
HSF方法名称：getListByCampusAndType
*/
type AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest struct {
    model.Params

    // 系统自动生成
    param0   *WorkBenchContext 

    // 查询参数封装
    param1   *SpaceGroupQuery 

}

func NewAlibabaCampusSpaceGroupGetlistbycampusandtypeRequest() *AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest{
    return &AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest) GetApiMethodName() string {
    return "alibaba.campus.space.group.getlistbycampusandtype"
}

func (r AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest) GetParam0() *WorkBenchContext {
    return r.param0
}

func (r *AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest) SetParam1(param1 *SpaceGroupQuery) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest) GetParam1() *SpaceGroupQuery {
    return r.param1
}

