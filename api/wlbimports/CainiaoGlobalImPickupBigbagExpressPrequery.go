package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupBigbagExpressPrequery 首公里揽收-快递预查询服务
// cainiao.global.im.pickup.bigbag.express.prequery
//
// 快递预查询服务
func CainiaoGlobalImPickupBigbagExpressPrequery(clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest, session string) (*wlbimports.CainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse, error) {
	var resp wlbimports.CainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
