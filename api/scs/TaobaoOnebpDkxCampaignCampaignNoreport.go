package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// Taobaoonebpdkxcampaigncampaignnoreport 获取场景计划的非报表数据
// taobao.onebp.dkx.campaign.campaign.noreport
//
// 获取场景计划的非报表数据。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
func Taobaoonebpdkxcampaigncampaignnoreport(clt *core.SDKClient, req *scs.TaobaoonebpdkxcampaigncampaignnoreportAPIRequest, session string) (*scs.TaobaoonebpdkxcampaigncampaignnoreportAPIResponse, error) {
	var resp scs.TaobaoonebpdkxcampaigncampaignnoreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
