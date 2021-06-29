package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据楼宇ID获取楼层 APIRequest
alibaba.campus.space.floor.getbybuildingid

根据楼宇ID获取楼层
HSF接口名称：com.alibaba.campus.api.space.service.top.FloorApiTopService
HSF方法名称：getFloorList
*/
type AlibabaCampusSpaceFloorGetbybuildingidRequest struct {
    model.Params

    // 系统自动生成
    param0   *WorkBenchContext 

    // 楼宇iD封装
    param1   *FloorQuery 

}

func NewAlibabaCampusSpaceFloorGetbybuildingidRequest() *AlibabaCampusSpaceFloorGetbybuildingidRequest{
    return &AlibabaCampusSpaceFloorGetbybuildingidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusSpaceFloorGetbybuildingidRequest) GetApiMethodName() string {
    return "alibaba.campus.space.floor.getbybuildingid"
}

func (r AlibabaCampusSpaceFloorGetbybuildingidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusSpaceFloorGetbybuildingidRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaCampusSpaceFloorGetbybuildingidRequest) GetParam0() *WorkBenchContext {
    return r.param0
}

func (r *AlibabaCampusSpaceFloorGetbybuildingidRequest) SetParam1(param1 *FloorQuery) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaCampusSpaceFloorGetbybuildingidRequest) GetParam1() *FloorQuery {
    return r.param1
}

