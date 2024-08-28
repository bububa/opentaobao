package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentTableauditOfflinelauncheritem 运营位管理-联盟一体机下线运营位内容
// yunos.tvpubadmin.content.tableaudit.offlinelauncheritem
//
// 运营位管理-联盟一体机下线运营位内容
func YunosTvpubadminContentTableauditOfflinelauncheritem(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest, resp *tvupadmin.YunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
