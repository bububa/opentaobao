package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCampaignPlatformUpdate 更新一个推广计划的平台设置
// taobao.simba.campaign.platform.update
//
// 更新一个推广计划的平台设置
func TaobaoSimbaCampaignPlatformUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignPlatformUpdateAPIRequest, resp *simba.TaobaoSimbaCampaignPlatformUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
