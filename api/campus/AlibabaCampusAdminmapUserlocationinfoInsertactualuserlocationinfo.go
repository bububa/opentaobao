package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfo 上传用户实时位置
// alibaba.campus.adminmap.userlocationinfo.insertactualuserlocationinfo
//
// 上传用户实时位置
// HSF接口名称：com.alibaba.campus.api.adminmap.service.top.UserLocationQueryApiTopService
// HSF方法名称：insertActualUserLocationInfo
func AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfo(clt *core.SDKClient, req *campus.AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest, resp *campus.AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
