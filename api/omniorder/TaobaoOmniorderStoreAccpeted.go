package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreAccpeted Pos端门店接单接口
// taobao.omniorder.store.accpeted
//
// ISV Pos端门店接单，通知星盘
func TaobaoOmniorderStoreAccpeted(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreAccpetedAPIRequest, resp *omniorder.TaobaoOmniorderStoreAccpetedAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
