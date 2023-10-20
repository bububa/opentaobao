package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCampaignTimetemplateFind 获取分时折扣模板
// taobao.simba.campaign.timetemplate.find
//
// 批量得到智能推广推广计划下的推广组
func TaobaoSimbaCampaignTimetemplateFind(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignTimetemplateFindAPIRequest, resp *simba.TaobaoSimbaCampaignTimetemplateFindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
