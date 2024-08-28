package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressOrderTmsUpdate 服务商修改上门取退时间接口
// taobao.logistics.express.order.tms.update
//
// 服务商修改上门取退时间接口
func TaobaoLogisticsExpressOrderTmsUpdate(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressOrderTmsUpdateAPIRequest, resp *logistic.TaobaoLogisticsExpressOrderTmsUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
