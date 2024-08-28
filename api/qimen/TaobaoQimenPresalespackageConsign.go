package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenPresalespackageConsign 预售预包尾款推单发货
// taobao.qimen.presalespackage.consign
//
// 预售预包尾款推单发货
func TaobaoQimenPresalespackageConsign(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenPresalespackageConsignAPIRequest, resp *qimen.TaobaoQimenPresalespackageConsignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
