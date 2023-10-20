package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceTvid 查询终端信息
// yunos.tvpubadmin.device.tvid
//
// 通过UUID查询终端信息
func YunosTvpubadminDeviceTvid(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceTvidAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceTvidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
