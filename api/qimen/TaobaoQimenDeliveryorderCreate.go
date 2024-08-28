package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenDeliveryorderCreate 发货单创建接口
// taobao.qimen.deliveryorder.create
//
// taobao.qimen.deliveryorder.create
func TaobaoQimenDeliveryorderCreate(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenDeliveryorderCreateAPIRequest, resp *qimen.TaobaoQimenDeliveryorderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
