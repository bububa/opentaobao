package openim

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeCreate 创建群
// taobao.openim.tribe.create
//
// 创建一个openim的群
func TaobaoOpenimTribeCreate(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimTribeCreateAPIRequest, resp *openim.TaobaoOpenimTribeCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
