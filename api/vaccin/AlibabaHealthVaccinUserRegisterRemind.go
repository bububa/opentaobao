package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinUserRegisterRemind isv到苗提醒
// alibaba.health.vaccin.user.register.remind
//
// isv到苗提醒
func AlibabaHealthVaccinUserRegisterRemind(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinUserRegisterRemindAPIRequest, session string) (*vaccin.AlibabaHealthVaccinUserRegisterRemindAPIResponse, error) {
	var resp vaccin.AlibabaHealthVaccinUserRegisterRemindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
