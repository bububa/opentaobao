package openim

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeSetmanager OPENIM群设置管理员
// taobao.openim.tribe.setmanager
//
// OPENIM群设置管理员
func TaobaoOpenimTribeSetmanager(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimTribeSetmanagerAPIRequest, resp *openim.TaobaoOpenimTribeSetmanagerAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
