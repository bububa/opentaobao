package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenTransferorderCreate 调拨单创建
// taobao.qimen.transferorder.create
//
// 调拨单创建
func TaobaoQimenTransferorderCreate(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenTransferorderCreateAPIRequest, resp *qimen.TaobaoQimenTransferorderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
