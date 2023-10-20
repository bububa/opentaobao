package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabahealthvaccinuserregisterremind isv到苗提醒
// alibaba.health.vaccin.user.register.remind
//
// isv到苗提醒
func Alibabahealthvaccinuserregisterremind(clt *core.SDKClient, req *vaccin.AlibabahealthvaccinuserregisterremindAPIRequest, session string) (*vaccin.AlibabahealthvaccinuserregisterremindAPIResponse, error) {
	var resp vaccin.AlibabahealthvaccinuserregisterremindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
