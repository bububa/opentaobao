package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceApkinfo 获取停开服apk信息
// yunos.tvpubadmin.device.apkinfo
//
// 获取停开服apk信息
func YunosTvpubadminDeviceApkinfo(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceApkinfoAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceApkinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
