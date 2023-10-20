package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacampaignsget 取得一组推广计划
// taobao.simba.campaigns.get
//
// 取得一个客户的推广计划；
func Taobaosimbacampaignsget(clt *core.SDKClient, req *simba.TaobaosimbacampaignsgetAPIRequest, session string) (*simba.TaobaosimbacampaignsgetAPIResponse, error) {
	var resp simba.TaobaosimbacampaignsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
