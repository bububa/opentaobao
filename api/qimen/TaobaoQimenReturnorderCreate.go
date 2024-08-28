package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenReturnorderCreate 退货入库单创建接口
// taobao.qimen.returnorder.create
//
// taobao.qimen.returnorder.create
func TaobaoQimenReturnorderCreate(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenReturnorderCreateAPIRequest, resp *qimen.TaobaoQimenReturnorderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
