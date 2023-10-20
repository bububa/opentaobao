package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceOsupgradequery 系统升级查询
// yunos.tvpubadmin.device.osupgradequery
//
// 系统升级查询
func YunosTvpubadminDeviceOsupgradequery(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceOsupgradequeryAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceOsupgradequeryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
