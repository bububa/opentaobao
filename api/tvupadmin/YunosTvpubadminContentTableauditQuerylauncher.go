package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentTableauditQuerylauncher 运营位管控-查询联盟一体机运营位元数据列表
// yunos.tvpubadmin.content.tableaudit.querylauncher
//
// 运营位管控-查询联盟一体机运营位元数据列表
func YunosTvpubadminContentTableauditQuerylauncher(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentTableauditQuerylauncherAPIRequest, resp *tvupadmin.YunosTvpubadminContentTableauditQuerylauncherAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
