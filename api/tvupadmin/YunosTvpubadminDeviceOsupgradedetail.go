package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceOsupgradedetail 获取系统升级详情
// yunos.tvpubadmin.device.osupgradedetail
//
// 获取系统升级详情
func YunosTvpubadminDeviceOsupgradedetail(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceOsupgradedetailAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceOsupgradedetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
