package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceAppupgradequery 应用升级查询
// yunos.tvpubadmin.device.appupgradequery
//
// 应用升级查询
func YunosTvpubadminDeviceAppupgradequery(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceAppupgradequeryAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceAppupgradequeryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
