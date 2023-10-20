package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServicestoreSave 保存网点
// alibaba.ssc.supplyplatform.servicestore.save
//
// 网点创建、修改
func AlibabaSscSupplyplatformServicestoreSave(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServicestoreSaveAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServicestoreSaveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
