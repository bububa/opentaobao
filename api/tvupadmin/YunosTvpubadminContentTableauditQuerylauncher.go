package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentTableauditQuerylauncher 运营位管控-查询联盟一体机运营位元数据列表
// yunos.tvpubadmin.content.tableaudit.querylauncher
//
// 运营位管控-查询联盟一体机运营位元数据列表
func YunosTvpubadminContentTableauditQuerylauncher(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentTableauditQuerylauncherAPIRequest, resp *tvupadmin.YunosTvpubadminContentTableauditQuerylauncherAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
