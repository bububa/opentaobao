package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceabilitySave 保存服务能力
// alibaba.ssc.supplyplatform.serviceability.save
//
// 保存服务能力
func AlibabaSscSupplyplatformServiceabilitySave(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceabilitySaveAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServiceabilitySaveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
