package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologs 分时间段获取用户历史位置信息
// alibaba.campus.adminmap.userlocationinfo.getuserlocationinfologs
//
// 分时间段获取用户历史位置信息
func AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologs(clt *core.SDKClient, req *campus.AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest, resp *campus.AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
