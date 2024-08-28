package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentChildRootnodeGet 获取少儿大厅根类目接口
// yunos.tvpubadmin.content.child.rootnode.get
//
// 通过此接口可获取少儿大厅根类目列表
func YunosTvpubadminContentChildRootnodeGet(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChildRootnodeGetAPIRequest, resp *tvupadmin.YunosTvpubadminContentChildRootnodeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
