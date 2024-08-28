package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenDeliveryorderBatchconfirm 发货单确认接口
// taobao.qimen.deliveryorder.batchconfirm
//
// taobao.qimen.deliveryorder.batchconfirm
func TaobaoQimenDeliveryorderBatchconfirm(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenDeliveryorderBatchconfirmAPIRequest, resp *qimen.TaobaoQimenDeliveryorderBatchconfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
