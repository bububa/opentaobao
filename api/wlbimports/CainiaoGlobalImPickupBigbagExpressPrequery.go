package wlbimports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupBigbagExpressPrequery 首公里揽收-快递预查询服务
// cainiao.global.im.pickup.bigbag.express.prequery
//
// 快递预查询服务
func CainiaoGlobalImPickupBigbagExpressPrequery(ctx context.Context, clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest, resp *wlbimports.CainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
