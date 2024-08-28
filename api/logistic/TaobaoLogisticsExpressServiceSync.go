package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressServiceSync 服务信息回告接口
// taobao.logistics.express.service.sync
//
// 服务信息回告接口
func TaobaoLogisticsExpressServiceSync(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressServiceSyncAPIRequest, resp *logistic.TaobaoLogisticsExpressServiceSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
