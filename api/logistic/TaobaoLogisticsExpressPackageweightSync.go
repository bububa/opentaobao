package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressPackageweightSync TMS包裹重量回传
// taobao.logistics.express.packageweight.sync
//
// TMS包裹重量回传
func TaobaoLogisticsExpressPackageweightSync(clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressPackageweightSyncAPIRequest, session string) (*logistic.TaobaoLogisticsExpressPackageweightSyncAPIResponse, error) {
	var resp logistic.TaobaoLogisticsExpressPackageweightSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
