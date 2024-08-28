package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressOrderPayTmsQuery 上门取退运费支付状态查询接口
// taobao.logistics.express.order.pay.tms.query
//
// 上门取退运费支付状态查询接口
func TaobaoLogisticsExpressOrderPayTmsQuery(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest, resp *logistic.TaobaoLogisticsExpressOrderPayTmsQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
