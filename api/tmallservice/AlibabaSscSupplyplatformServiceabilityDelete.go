package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

/* AlibabaSscSupplyplatformServiceabilityDelete
删除服务能力
alibaba.ssc.supplyplatform.serviceability.delete

删除服务能力 */
func AlibabaSscSupplyplatformServiceabilityDelete(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest, session string) (*tmallservice.AlibabaSscSupplyplatformServiceabilityDeleteAPIResponse, error) {
	var resp tmallservice.AlibabaSscSupplyplatformServiceabilityDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
