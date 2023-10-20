package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceworkerSave 服务商绑定工人
// alibaba.ssc.supplyplatform.serviceworker.save
//
// 服务商将上传工人与服务商自己建立关系，需要将工人的服务区域和住址回传
func AlibabaSscSupplyplatformServiceworkerSave(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceworkerSaveAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServiceworkerSaveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
