package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinSubscribeInfoReturn 自有pov预约信息回传
// alibaba.health.vaccin.subscribe.info.return
//
// 自有pov预约信息回传
func AlibabaHealthVaccinSubscribeInfoReturn(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinSubscribeInfoReturnAPIRequest, session string) (*vaccin.AlibabaHealthVaccinSubscribeInfoReturnAPIResponse, error) {
	var resp vaccin.AlibabaHealthVaccinSubscribeInfoReturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
