package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinUserRegisterRemind isv到苗提醒
// alibaba.health.vaccin.user.register.remind
//
// isv到苗提醒
func AlibabaHealthVaccinUserRegisterRemind(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinUserRegisterRemindAPIRequest, resp *vaccin.AlibabaHealthVaccinUserRegisterRemindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
