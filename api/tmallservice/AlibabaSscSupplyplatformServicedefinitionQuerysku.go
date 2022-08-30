package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServicedefinitionQuerysku 服务sku查询
// alibaba.ssc.supplyplatform.servicedefinition.querysku
//
// 服务sku查询
func AlibabaSscSupplyplatformServicedefinitionQuerysku(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest, session string) (*tmallservice.AlibabaSscSupplyplatformServicedefinitionQueryskuAPIResponse, error) {
	var resp tmallservice.AlibabaSscSupplyplatformServicedefinitionQueryskuAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
