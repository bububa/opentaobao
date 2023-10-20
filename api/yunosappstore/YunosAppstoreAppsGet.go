package yunosappstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosappstore"
)

// YunosAppstoreAppsGet 根据包名列表获取应用信息列表
// yunos.appstore.apps.get
//
// 根据包名列表获取应用信息列表
func YunosAppstoreAppsGet(clt *core.SDKClient, req *yunosappstore.YunosAppstoreAppsGetAPIRequest, resp *yunosappstore.YunosAppstoreAppsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
