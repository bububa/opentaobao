package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDiccontroltaskQuery 停开服任务列表
// yunos.tvpubadmin.diccontroltask.query
//
// 牌照方对终端设备的停开服管理
func YunosTvpubadminDiccontroltaskQuery(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDiccontroltaskQueryAPIRequest, resp *tvupadmin.YunosTvpubadminDiccontroltaskQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
