package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据分组ID查询相应的空间单元 API请求
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

// 初始化AlibabaCampusSpaceUnitGetlistbygroupidRequest对象
func NewAlibabaCampusSpaceUnitGetlistbygroupidRequest() *AlibabaCampusSpaceUnitGetlistbygroupidRequest{
    return &AlibabaCampusSpaceUnitGetlistbygroupidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceUnitGetlistbygroupidRequest) GetApiMethodName() string {
    return "alibaba.campus.space.unit.getlistbygroupid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceUnitGetlistbygroupidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 分组ID
func (r *AlibabaCampusSpaceUnitGetlistbygroupidRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusSpaceUnitGetlistbygroupidRequest) GetParam0() *WorkBenchContext {
    return r.param0
}
// Param1 Setter
// 分组ID
func (r *AlibabaCampusSpaceUnitGetlistbygroupidRequest) SetParam1(param1 int64) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusSpaceUnitGetlistbygroupidRequest) GetParam1() int64 {
    return r.param1
}
