package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpcampaigngroupfindlist 查询计划组列表
// taobao.universalbp.campaigngroup.findlist
//
// 查询某个场景内的计划组列表
func Taobaouniversalbpcampaigngroupfindlist(clt *core.SDKClient, req *simba.TaobaouniversalbpcampaigngroupfindlistAPIRequest, session string) (*simba.TaobaouniversalbpcampaigngroupfindlistAPIResponse, error) {
	var resp simba.TaobaouniversalbpcampaigngroupfindlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
