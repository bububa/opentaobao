package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressPackageweightSync TMS包裹重量回传
// taobao.logistics.express.packageweight.sync
//
// TMS包裹重量回传
func TaobaoLogisticsExpressPackageweightSync(clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressPackageweightSyncAPIRequest, resp *logistic.TaobaoLogisticsExpressPackageweightSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
