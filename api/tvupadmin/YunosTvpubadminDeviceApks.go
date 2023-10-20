package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceApks 获取停开服apk列表
// yunos.tvpubadmin.device.apks
//
// 获取停开服apk列表
func YunosTvpubadminDeviceApks(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceApksAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceApksAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
