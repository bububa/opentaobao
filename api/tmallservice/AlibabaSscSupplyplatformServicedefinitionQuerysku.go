package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServicedefinitionQuerysku 服务sku查询
// alibaba.ssc.supplyplatform.servicedefinition.querysku
//
// 服务sku查询
func AlibabaSscSupplyplatformServicedefinitionQuerysku(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServicedefinitionQueryskuAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
