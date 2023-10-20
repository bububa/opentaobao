package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressDeliveryCutNotify TMS配拦截结果回告
// taobao.logistics.express.delivery.cut.notify
//
// TMS配拦截结果回告
func TaobaoLogisticsExpressDeliveryCutNotify(clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressDeliveryCutNotifyAPIRequest, resp *logistic.TaobaoLogisticsExpressDeliveryCutNotifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
