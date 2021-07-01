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
type AlibabaCampusSpaceGroupGetlistAPIRequest struct {
    model.Params
    // 查询条件封装
    _param0   *WorkBenchContext
    // 查询参数封装
    _param1   *SpaceGroupQuery
}

// 初始化AlibabaCampusSpaceGroupGetlistAPIRequest对象
func NewAlibabaCampusSpaceGroupGetlistRequest() *AlibabaCampusSpaceGroupGetlistAPIRequest{
    return &AlibabaCampusSpaceGroupGetlistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceGroupGetlistAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.space.group.getlist"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceGroupGetlistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 查询条件封装
func (r *AlibabaCampusSpaceGroupGetlistAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusSpaceGroupGetlistAPIRequest) GetParam0() *WorkBenchContext {
    return r._param0
}
// Param1 Setter
// 查询参数封装
func (r *AlibabaCampusSpaceGroupGetlistAPIRequest) SetParam1(_param1 *SpaceGroupQuery) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusSpaceGroupGetlistAPIRequest) GetParam1() *SpaceGroupQuery {
    return r._param1
}
