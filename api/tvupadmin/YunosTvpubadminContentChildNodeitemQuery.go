package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentChildNodeitemQuery 查询少儿大厅类目内容
// yunos.tvpubadmin.content.child.nodeitem.query
//
// 查询少儿大厅类目内容信息
func YunosTvpubadminContentChildNodeitemQuery(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChildNodeitemQueryAPIRequest, resp *tvupadmin.YunosTvpubadminContentChildNodeitemQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
