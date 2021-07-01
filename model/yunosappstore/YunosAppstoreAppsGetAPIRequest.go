package yunosappstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosAppstoreAppsGetAPIRequest
根据包名列表获取应用信息列表 API请求
yunos.appstore.apps.get

根据包名列表获取应用信息列表 */
type YunosAppstoreAppsGetAPIRequest struct {
	model.Params
	// 应用包名列表
	_pkgs []string
}

// New
