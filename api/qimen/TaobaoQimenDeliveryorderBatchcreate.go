package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenDeliveryorderBatchcreate 发货单创建批量接口
// taobao.qimen.deliveryorder.batchcreate
//
// taobao.qimen.deliveryorder.batchcreate
func TaobaoQimenDeliveryorderBatchcreate(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenDeliveryorderBatchcreateAPIRequest, resp *qimen.TaobaoQimenDeliveryorderBatchcreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
