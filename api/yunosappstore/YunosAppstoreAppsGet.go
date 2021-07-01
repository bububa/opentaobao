package yunosappstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosappstore"
)

/* YunosAppstoreAppsGet
根据包名列表获取应用信息列表
yunos.appstore.apps.get

根据包名列表获取应用信息列表 */
func YunosAppstoreAppsGet(clt *core.SDKClient, req *yunosappstore.YunosAppstoreAppsGetAPIRequest, session string) (*yunosappstore.YunosAppstoreAppsGetAPIResponse, error) {
	var resp yunosappstore.YunosAppstoreAppsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
