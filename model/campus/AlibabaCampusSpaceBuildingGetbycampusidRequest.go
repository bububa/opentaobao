package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据园区ID获取楼宇 APIRequest
alibaba.campus.space.building.getbycampusid

根据园区ID获取楼宇
HSF接口名称：com.alibaba.campus.api.space.service.top.BuildingApiTopService
HSF方法名称：getBuildingList
*/
type AlibabaCampusSpaceBuildingGetbycampusidRequest struct {
    model.Params

    // 系统自动生成
    param0   *WorkBenchContext 

    // 园区封装
    param1   *BuildingQuery 

}

func NewAlibabaCampusSpaceBuildingGetbycampusidRequest() *AlibabaCampusSpaceBuildingGetbycampusidRequest{
    return &AlibabaCampusSpaceBuildingGetbycampusidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusSpaceBuildingGetbycampusidRequest) GetApiMethodName() string {
    return "alibaba.campus.space.building.getbycampusid"
}

func (r AlibabaCampusSpaceBuildingGetbycampusidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusSpaceBuildingGetbycampusidRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaCampusSpaceBuildingGetbycampusidRequest) GetParam0() *WorkBenchContext {
    return r.param0
}

func (r *AlibabaCampusSpaceBuildingGetbycampusidRequest) SetParam1(param1 *BuildingQuery) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaCampusSpaceBuildingGetbycampusidRequest) GetParam1() *BuildingQuery {
    return r.param1
}

