package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据楼宇ID获取楼层 API请求
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

// 初始化AlibabaCampusSpaceFloorGetbybuildingidRequest对象
func NewAlibabaCampusSpaceFloorGetbybuildingidRequest() *AlibabaCampusSpaceFloorGetbybuildingidRequest{
    return &AlibabaCampusSpaceFloorGetbybuildingidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceFloorGetbybuildingidRequest) GetApiMethodName() string {
    return "alibaba.campus.space.floor.getbybuildingid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceFloorGetbybuildingidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 系统自动生成
func (r *AlibabaCampusSpaceFloorGetbybuildingidRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusSpaceFloorGetbybuildingidRequest) GetParam0() *WorkBenchContext {
    return r.param0
}
// Param1 Setter
// 楼宇iD封装
func (r *AlibabaCampusSpaceFloorGetbybuildingidRequest) SetParam1(param1 *FloorQuery) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusSpaceFloorGetbybuildingidRequest) GetParam1() *FloorQuery {
    return r.param1
}
