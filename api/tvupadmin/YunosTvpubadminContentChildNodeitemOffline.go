package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentChildNodeitemOffline 少儿大厅类目内容下线接口
// yunos.tvpubadmin.content.child.nodeitem.offline
//
// 少儿大厅类目内容下线接口
func YunosTvpubadminContentChildNodeitemOffline(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChildNodeitemOfflineAPIRequest, resp *tvupadmin.YunosTvpubadminContentChildNodeitemOfflineAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
