package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpCampaigngroupFindlist 查询计划组列表
// taobao.universalbp.campaigngroup.findlist
//
// 查询某个场景内的计划组列表
func TaobaoUniversalbpCampaigngroupFindlist(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpCampaigngroupFindlistAPIRequest, resp *simba.TaobaoUniversalbpCampaigngroupFindlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
