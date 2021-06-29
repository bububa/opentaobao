package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
多条件查询空间分组信息 API请求
alibaba.campus.space.group.getlist

多条件查询空间分组信息
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
HSF方法名称：getList
*/
type AlibabaCampusSpaceGroupGetlistRequest struct {
    model.Params
    // 查询条件封装
    param0   *WorkBenchContext
    // 查询参数封装
    param1   *SpaceGroupQuery
}

// 初始化AlibabaCampusSpaceGroupGetlistRequest对象
func NewAlibabaCampusSpaceGroupGetlistRequest() *AlibabaCampusSpaceGroupGetlistRequest{
    return &AlibabaCampusSpaceGroupGetlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceGroupGetlistRequest) GetApiMethodName() string {
    return "alibaba.campus.space.group.getlist"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceGroupGetlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 查询条件封装
func (r *AlibabaCampusSpaceGroupGetlistRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusSpaceGroupGetlistRequest) GetParam0() *WorkBenchContext {
    return r.param0
}
// Param1 Setter
// 查询参数封装
func (r *AlibabaCampusSpaceGroupGetlistRequest) SetParam1(param1 *SpaceGroupQuery) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusSpaceGroupGetlistRequest) GetParam1() *SpaceGroupQuery {
    return r.param1
}
