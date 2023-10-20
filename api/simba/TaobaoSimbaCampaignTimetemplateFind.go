package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacampaigntimetemplatefind 获取分时折扣模板
// taobao.simba.campaign.timetemplate.find
//
// 批量得到智能推广推广计划下的推广组
func Taobaosimbacampaigntimetemplatefind(clt *core.SDKClient, req *simba.TaobaosimbacampaigntimetemplatefindAPIRequest, session string) (*simba.TaobaosimbacampaigntimetemplatefindAPIResponse, error) {
	var resp simba.TaobaosimbacampaigntimetemplatefindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
