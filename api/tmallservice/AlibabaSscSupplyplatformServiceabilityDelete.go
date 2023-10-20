package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceabilityDelete 删除服务能力
// alibaba.ssc.supplyplatform.serviceability.delete
//
// 删除服务能力
func AlibabaSscSupplyplatformServiceabilityDelete(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServiceabilityDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
