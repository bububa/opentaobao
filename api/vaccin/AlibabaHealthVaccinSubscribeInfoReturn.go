package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabahealthvaccinsubscribeinforeturn 自有pov预约信息回传
// alibaba.health.vaccin.subscribe.info.return
//
// 自有pov预约信息回传
func Alibabahealthvaccinsubscribeinforeturn(clt *core.SDKClient, req *vaccin.AlibabahealthvaccinsubscribeinforeturnAPIRequest, session string) (*vaccin.AlibabahealthvaccinsubscribeinforeturnAPIResponse, error) {
	var resp vaccin.AlibabahealthvaccinsubscribeinforeturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
