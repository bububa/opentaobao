package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
根据userId(支持单个或批量)获取用户实时位置信息 
alibaba.campus.adminmap.userlocationinfo.getactualuserlocationinfobyids

根据userId(支持单个或批量)获取用户实时位置信息 
HSF接口名称：com.alibaba.campus.api.adminmap.service.top.UserLocationQueryApiTopService
HSF方法名称：getActualUserLocationInfoByIds
*/
func AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyids(clt *core.SDKClient, req *campus.AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsRequest, session string) (*campus.AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse, error) {
    var resp campus.AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
