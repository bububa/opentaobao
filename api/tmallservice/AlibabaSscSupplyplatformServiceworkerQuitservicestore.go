package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceworkerQuitservicestore 工人退出网点
// alibaba.ssc.supplyplatform.serviceworker.quitservicestore
//
// 工人退出网点
func AlibabaSscSupplyplatformServiceworkerQuitservicestore(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
