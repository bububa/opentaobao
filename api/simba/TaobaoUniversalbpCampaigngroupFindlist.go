package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpCampaigngroupFindlist 查询计划组列表
// taobao.universalbp.campaigngroup.findlist
//
// 查询某个场景内的计划组列表
func TaobaoUniversalbpCampaigngroupFindlist(clt *core.SDKClient, req *simba.TaobaoUniversalbpCampaigngroupFindlistAPIRequest, session string) (*simba.TaobaoUniversalbpCampaigngroupFindlistAPIResponse, error) {
	var resp simba.TaobaoUniversalbpCampaigngroupFindlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
