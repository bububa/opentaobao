package openim

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeQuit OPENIM群成员退出
// taobao.openim.tribe.quit
//
// OPENIM群成员退出
func TaobaoOpenimTribeQuit(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimTribeQuitAPIRequest, resp *openim.TaobaoOpenimTribeQuitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
