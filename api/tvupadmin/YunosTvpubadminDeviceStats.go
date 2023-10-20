package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceStats 获取设备统计数据
// yunos.tvpubadmin.device.stats
//
// 获取设备统计数据
func YunosTvpubadminDeviceStats(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceStatsAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceStatsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
