package openim

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeJoin OPENIM群主动加入
// taobao.openim.tribe.join
//
// OPENIM群主动加入
func TaobaoOpenimTribeJoin(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimTribeJoinAPIRequest, resp *openim.TaobaoOpenimTribeJoinAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
