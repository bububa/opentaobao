package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据userId(支持单个或批量)获取用户实时位置信息 API请求
alibaba.campus.adminmap.userlocationinfo.getactualuserlocationinfobyids

根据userId(支持单个或批量)获取用户实时位置信息 
HSF接口名称：com.alibaba.campus.api.adminmap.service.top.UserLocationQueryApiTopService
HSF方法名称：getActualUserLocationInfoByIds
*/
type AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsRequest struct {
    model.Params
    // 环境参数
    param0   *WorkBenchContext
    // 查询参数
    param1   *UserLocationInfoQuery
}

// 初始化AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsRequest对象
func NewAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsRequest() *AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsRequest{
    return &AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsRequest) GetApiMethodName() string {
    return "alibaba.campus.adminmap.userlocationinfo.getactualuserlocationinfobyids"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 环境参数
func (r *AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsRequest) GetParam0() *WorkBenchContext {
    return r.param0
}
// Param1 Setter
// 查询参数
func (r *AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsRequest) SetParam1(param1 *UserLocationInfoQuery) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsRequest) GetParam1() *UserLocationInfoQuery {
    return r.param1
}
