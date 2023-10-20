package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceQuery 获取设备列表
// yunos.tvpubadmin.device.query
//
// 获取设备列表
func YunosTvpubadminDeviceQuery(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceQueryAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
