package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceworkerCancelleave 工人取消请假
// alibaba.ssc.supplyplatform.serviceworker.cancelleave
//
// 工人取消请假
func AlibabaSscSupplyplatformServiceworkerCancelleave(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
