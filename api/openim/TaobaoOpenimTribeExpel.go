package openim

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeExpel OPENIM群踢出成员
// taobao.openim.tribe.expel
//
// OPENIM群踢出成员
func TaobaoOpenimTribeExpel(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimTribeExpelAPIRequest, resp *openim.TaobaoOpenimTribeExpelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
