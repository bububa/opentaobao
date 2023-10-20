package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressOrderPayTmsQuery 上门取退运费支付状态查询接口
// taobao.logistics.express.order.pay.tms.query
//
// 上门取退运费支付状态查询接口
func TaobaoLogisticsExpressOrderPayTmsQuery(clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest, resp *logistic.TaobaoLogisticsExpressOrderPayTmsQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
