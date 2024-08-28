package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenOrderPending 单据挂起（恢复）接口
// taobao.qimen.order.pending
//
// ERP调用奇门的接口,挂起某些创建的单据;场景介绍：ERP主动发起挂起（恢复）某些创建的单据，如入库单、出库单、退货单等
func TaobaoQimenOrderPending(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenOrderPendingAPIRequest, resp *qimen.TaobaoQimenOrderPendingAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
