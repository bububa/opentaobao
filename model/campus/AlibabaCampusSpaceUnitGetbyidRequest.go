package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据ID查询指定空间单元信息 API请求
alibaba.campus.space.unit.getbyid

根据ID查询指定空间单元信息
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：getById
*/
type AlibabaCampusSpaceUnitGetbyidAPIRequest struct {
    model.Params
    // 空间单元ID
    _param0   *WorkBenchContext
    // 空间单元ID
    _param1   int64
}

// 初始化AlibabaCampusSpaceUnitGetbyidAPIRequest对象
func NewAlibabaCampusSpaceUnitGetbyidRequest() *AlibabaCampusSpaceUnitGetbyidAPIRequest{
    return &AlibabaCampusSpaceUnitGetbyidAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceUnitGetbyidAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.space.unit.getbyid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceUnitGetbyidAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 空间单元ID
func (r *AlibabaCampusSpaceUnitGetbyidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusSpaceUnitGetbyidAPIRequest) GetParam0() *WorkBenchContext {
    return r._param0
}
// Param1 Setter
// 空间单元ID
func (r *AlibabaCampusSpaceUnitGetbyidAPIRequest) SetParam1(_param1 int64) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusSpaceUnitGetbyidAPIRequest) GetParam1() int64 {
    return r._param1
}
