package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreRefused Pos端门店拒单
// taobao.omniorder.store.refused
//
// ISV Pos端门店拒单，通知星盘
func TaobaoOmniorderStoreRefused(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreRefusedAPIRequest, resp *omniorder.TaobaoOmniorderStoreRefusedAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
