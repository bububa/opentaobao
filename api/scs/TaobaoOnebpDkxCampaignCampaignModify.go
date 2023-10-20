package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// Taobaoonebpdkxcampaigncampaignmodify 修改计划
// taobao.onebp.dkx.campaign.campaign.modify
//
// 修改计划。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
func Taobaoonebpdkxcampaigncampaignmodify(clt *core.SDKClient, req *scs.TaobaoonebpdkxcampaigncampaignmodifyAPIRequest, session string) (*scs.TaobaoonebpdkxcampaigncampaignmodifyAPIResponse, error) {
	var resp scs.TaobaoonebpdkxcampaigncampaignmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
