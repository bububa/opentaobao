package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// TaobaoLogisticsAppletPackageQuery 淘宝包裹查询
// taobao.logistics.applet.package.query
//
// 淘宝包裹查询
func TaobaoLogisticsAppletPackageQuery(clt *core.SDKClient, req *mtopopen.TaobaoLogisticsAppletPackageQueryAPIRequest, resp *mtopopen.TaobaoLogisticsAppletPackageQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
