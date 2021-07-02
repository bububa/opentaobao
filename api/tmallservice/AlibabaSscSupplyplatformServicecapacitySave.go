package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServicecapacitySave 保存服务容量
// alibaba.ssc.supplyplatform.servicecapacity.save
//
// 保存服务容量
func AlibabaSscSupplyplatformServicecapacitySave(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServicecapacitySaveAPIRequest, session string) (*tmallservice.AlibabaSscSupplyplatformServicecapacitySaveAPIResponse, error) {
	var resp tmallservice.AlibabaSscSupplyplatformServicecapacitySaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
