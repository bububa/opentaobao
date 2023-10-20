package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServicecapacityDelete 服务容量删除
// alibaba.ssc.supplyplatform.servicecapacity.delete
//
// 服务容量删除
func AlibabaSscSupplyplatformServicecapacityDelete(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServicecapacityDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
