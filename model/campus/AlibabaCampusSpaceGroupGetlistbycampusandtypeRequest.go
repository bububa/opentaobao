package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据园区id及TypeId获取空间分组 API请求
alibaba.campus.space.group.getlistbycampusandtype

根据园区id及TypeId获取空间分组
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
HSF方法名称：getListByCampusAndType
*/
type AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest struct {
    model.Params
    // 系统自动生成
    _param0   *WorkBenchContext
    // 查询参数封装
    _param1   *SpaceGroupQuery
}

// 初始化AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest对象
func NewAlibabaCampusSpaceGroupGetlistbycampusandtypeRequest() *AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest{
    return &AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest) GetApiMethodName() string {
    return "alibaba.campus.space.group.getlistbycampusandtype"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 系统自动生成
func (r *AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest) SetParam0(_param0 *WorkBenchContext) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest) GetParam0() *WorkBenchContext {
    return r._param0
}
// Param1 Setter
// 查询参数封装
func (r *AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest) SetParam1(_param1 *SpaceGroupQuery) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusSpaceGroupGetlistbycampusandtypeRequest) GetParam1() *SpaceGroupQuery {
    return r._param1
}
