package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceworkerQuitservicestore 工人退出网点
// alibaba.ssc.supplyplatform.serviceworker.quitservicestore
//
// 工人退出网点
func AlibabaSscSupplyplatformServiceworkerQuitservicestore(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
