package openim

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeGetmembers OPENIM群成员获取
// taobao.openim.tribe.getmembers
//
// OPENIM群成员获取
func TaobaoOpenimTribeGetmembers(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimTribeGetmembersAPIRequest, resp *openim.TaobaoOpenimTribeGetmembersAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
