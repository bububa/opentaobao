package openim

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeDismiss OPENIM群解散
// taobao.openim.tribe.dismiss
//
// OPENIM群解散
func TaobaoOpenimTribeDismiss(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimTribeDismissAPIRequest, resp *openim.TaobaoOpenimTribeDismissAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
