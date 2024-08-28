package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentChildLeafnodeGet 获取少儿大厅二级类目列表
// yunos.tvpubadmin.content.child.leafnode.get
//
// 获取少儿大厅二级类目列表的接口
func YunosTvpubadminContentChildLeafnodeGet(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChildLeafnodeGetAPIRequest, resp *tvupadmin.YunosTvpubadminContentChildLeafnodeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
