package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分时间段获取用户历史位置信息 API请求
alibaba.campus.adminmap.userlocationinfo.getuserlocationinfologs

分时间段获取用户历史位置信息
*/
type AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsRequest struct {
    model.Params
    // 环境参数
    _param0   *WorkBenchContext
    // 查询参数
    _param1   *UserLocationInfoQuery
}

// 初始化AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsRequest对象
func NewAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsRequest() *AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsRequest{
    return &AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsRequest) GetApiMethodName() string {
    return "alibaba.campus.adminmap.userlocationinfo.getuserlocationinfologs"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 环境参数
func (r *AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsRequest) SetParam0(_param0 *WorkBenchContext) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsRequest) GetParam0() *WorkBenchContext {
    return r._param0
}
// Param1 Setter
// 查询参数
func (r *AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsRequest) SetParam1(_param1 *UserLocationInfoQuery) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsRequest) GetParam1() *UserLocationInfoQuery {
    return r._param1
}
