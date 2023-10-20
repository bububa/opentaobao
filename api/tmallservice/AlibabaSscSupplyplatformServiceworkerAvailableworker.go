package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceworkerAvailableworker 查询可用工人
// alibaba.ssc.supplyplatform.serviceworker.availableworker
//
// 可用工人查询
func AlibabaSscSupplyplatformServiceworkerAvailableworker(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest, session string) (*tmallservice.AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIResponse, error) {
	var resp tmallservice.AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
