package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServicecapacitySave 保存服务容量
// alibaba.ssc.supplyplatform.servicecapacity.save
//
// 保存服务容量
func AlibabaSscSupplyplatformServicecapacitySave(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServicecapacitySaveAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServicecapacitySaveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
