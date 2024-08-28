package openim

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeUnsetmanager OPENIM群取消管理员
// taobao.openim.tribe.unsetmanager
//
// OPENIM群取消管理员
func TaobaoOpenimTribeUnsetmanager(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimTribeUnsetmanagerAPIRequest, resp *openim.TaobaoOpenimTribeUnsetmanagerAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
