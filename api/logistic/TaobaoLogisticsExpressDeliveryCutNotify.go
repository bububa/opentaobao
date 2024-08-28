package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressDeliveryCutNotify TMS配拦截结果回告
// taobao.logistics.express.delivery.cut.notify
//
// TMS配拦截结果回告
func TaobaoLogisticsExpressDeliveryCutNotify(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressDeliveryCutNotifyAPIRequest, resp *logistic.TaobaoLogisticsExpressDeliveryCutNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
