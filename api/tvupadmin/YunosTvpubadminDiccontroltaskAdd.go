package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDiccontroltaskAdd 新增停开服任务
// yunos.tvpubadmin.diccontroltask.add
//
// 新增停开服任务
func YunosTvpubadminDiccontroltaskAdd(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDiccontroltaskAddAPIRequest, resp *tvupadmin.YunosTvpubadminDiccontroltaskAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
