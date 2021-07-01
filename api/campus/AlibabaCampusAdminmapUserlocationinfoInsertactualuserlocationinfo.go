package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
上传用户实时位置 
alibaba.campus.adminmap.userlocationinfo.insertactualuserlocationinfo

上传用户实时位置
HSF接口名称：com.alibaba.campus.api.adminmap.service.top.UserLocationQueryApiTopService
HSF方法名称：insertActualUserLocationInfo
*/
func AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfo(clt *core.SDKClient, req *campus.AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest, session string) (*campus.AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse, error) {
    var resp campus.AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
