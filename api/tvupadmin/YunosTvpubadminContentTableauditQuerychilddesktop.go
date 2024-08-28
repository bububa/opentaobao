package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentTableauditQuerychilddesktop 迎客松查看小酷宝桌面坑位元数据列表
// yunos.tvpubadmin.content.tableaudit.querychilddesktop
//
// 迎客松查看小酷宝桌面坑位元数据列表
func YunosTvpubadminContentTableauditQuerychilddesktop(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest, resp *tvupadmin.YunosTvpubadminContentTableauditQuerychilddesktopAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
