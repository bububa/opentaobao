package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCampaignScheduleGet 取得一个推广计划的分时折扣设置
// taobao.simba.campaign.schedule.get
//
// 取得一个推广计划的分时折扣设置
func TaobaoSimbaCampaignScheduleGet(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignScheduleGetAPIRequest, resp *simba.TaobaoSimbaCampaignScheduleGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
