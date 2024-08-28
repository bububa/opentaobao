package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentTableauditQuerymetaitem 运营位管控-查询魔盒运营位元数据列表
// yunos.tvpubadmin.content.tableaudit.querymetaitem
//
// 运营位管控-查询魔盒运营位元数据列表
func YunosTvpubadminContentTableauditQuerymetaitem(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentTableauditQuerymetaitemAPIRequest, resp *tvupadmin.YunosTvpubadminContentTableauditQuerymetaitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
