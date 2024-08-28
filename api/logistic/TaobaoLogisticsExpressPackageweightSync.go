package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressPackageweightSync TMS包裹重量回传
// taobao.logistics.express.packageweight.sync
//
// TMS包裹重量回传
func TaobaoLogisticsExpressPackageweightSync(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressPackageweightSyncAPIRequest, resp *logistic.TaobaoLogisticsExpressPackageweightSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
