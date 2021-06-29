package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传用户实时位置 APIRequest
alibaba.campus.adminmap.userlocationinfo.insertactualuserlocationinfo

上传用户实时位置
HSF接口名称：com.alibaba.campus.api.adminmap.service.top.UserLocationQueryApiTopService
HSF方法名称：insertActualUserLocationInfo
*/
type AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoRequest struct {
    model.Params

    // 环境参数
    param0   *WorkBenchContext 

    // 查询参数
    param1   *UserLocationInfo 

}

func NewAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoRequest() *AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoRequest{
    return &AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoRequest) GetApiMethodName() string {
    return "alibaba.campus.adminmap.userlocationinfo.insertactualuserlocationinfo"
}

func (r AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoRequest) GetParam0() *WorkBenchContext {
    return r.param0
}

func (r *AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoRequest) SetParam1(param1 *UserLocationInfo) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoRequest) GetParam1() *UserLocationInfo {
    return r.param1
}

