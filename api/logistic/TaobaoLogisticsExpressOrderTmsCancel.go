package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressOrderTmsCancel 服务商上门取退时间取消接口
// taobao.logistics.express.order.tms.cancel
//
// 服务商上门取退时间取消接口
func TaobaoLogisticsExpressOrderTmsCancel(clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressOrderTmsCancelAPIRequest, resp *logistic.TaobaoLogisticsExpressOrderTmsCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
