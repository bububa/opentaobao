package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceworkerQuitservicestore 工人退出网点
// alibaba.ssc.supplyplatform.serviceworker.quitservicestore
//
// 工人退出网点
func AlibabaSscSupplyplatformServiceworkerQuitservicestore(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest, session string) (*tmallservice.AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse, error) {
	var resp tmallservice.AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
