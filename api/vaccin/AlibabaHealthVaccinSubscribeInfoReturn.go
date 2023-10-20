package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinSubscribeInfoReturn 自有pov预约信息回传
// alibaba.health.vaccin.subscribe.info.return
//
// 自有pov预约信息回传
func AlibabaHealthVaccinSubscribeInfoReturn(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinSubscribeInfoReturnAPIRequest, resp *vaccin.AlibabaHealthVaccinSubscribeInfoReturnAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
