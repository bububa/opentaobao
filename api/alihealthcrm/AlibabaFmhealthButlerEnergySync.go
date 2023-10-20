package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaFmhealthButlerEnergySync 同步用户消耗能量
// alibaba.fmhealth.butler.energy.sync
//
// 同步用户消耗能量，用户消耗s点或卡路里后，同步给健康平台
func AlibabaFmhealthButlerEnergySync(clt *core.SDKClient, req *alihealthcrm.AlibabaFmhealthButlerEnergySyncAPIRequest, session string) (*alihealthcrm.AlibabaFmhealthButlerEnergySyncAPIResponse, error) {
	var resp alihealthcrm.AlibabaFmhealthButlerEnergySyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
