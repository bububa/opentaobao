package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceworkerCancelleave 工人取消请假
// alibaba.ssc.supplyplatform.serviceworker.cancelleave
//
// 工人取消请假
func AlibabaSscSupplyplatformServiceworkerCancelleave(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest, session string) (*tmallservice.AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse, error) {
	var resp tmallservice.AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
