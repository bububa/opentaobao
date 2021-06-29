package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据园区id及TypeId获取空间单元 API请求
alibaba.campus.space.unit.getlistbycampusandtype

根据园区id及TypeId获取空间单元
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：getListByCampusAndType
*/
type AlibabaCampusSpaceUnitGetlistbycampusandtypeRequest struct {
    model.Params
    // 系统自动生成
    param0   *WorkBenchContext
    // 查询参数封装
    param1   *SpaceUnitQuery
}

// 初始化AlibabaCampusSpaceUnitGetlistbycampusandtypeRequest对象
func NewAlibabaCampusSpaceUnitGetlistbycampusandtypeRequest() *AlibabaCampusSpaceUnitGetlistbycampusandtypeRequest{
    return &AlibabaCampusSpaceUnitGetlistbycampusandtypeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceUnitGetlistbycampusandtypeRequest) GetApiMethodName() string {
    return "alibaba.campus.space.unit.getlistbycampusandtype"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceUnitGetlistbycampusandtypeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 系统自动生成
func (r *AlibabaCampusSpaceUnitGetlistbycampusandtypeRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusSpaceUnitGetlistbycampusandtypeRequest) GetParam0() *WorkBenchContext {
    return r.param0
}
// Param1 Setter
// 查询参数封装
func (r *AlibabaCampusSpaceUnitGetlistbycampusandtypeRequest) SetParam1(param1 *SpaceUnitQuery) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusSpaceUnitGetlistbycampusandtypeRequest) GetParam1() *SpaceUnitQuery {
    return r.param1
}
