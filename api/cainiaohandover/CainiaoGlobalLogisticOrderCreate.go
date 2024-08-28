package cainiaohandover

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalLogisticOrderCreate 创建物流订单
// cainiao.global.logistic.order.create
//
// 创建物流订单
func CainiaoGlobalLogisticOrderCreate(ctx context.Context, clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalLogisticOrderCreateAPIRequest, resp *cainiaohandover.CainiaoGlobalLogisticOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
