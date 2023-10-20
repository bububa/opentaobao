package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceYksBots 获取设备列表
// yunos.tvpubadmin.device.yks.bots
//
// 获取设备列表
func YunosTvpubadminDeviceYksBots(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceYksBotsAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceYksBotsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
