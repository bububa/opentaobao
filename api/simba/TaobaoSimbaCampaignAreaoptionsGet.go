package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacampaignareaoptionsget 取得推广计划的可设置投放地域列表
// taobao.simba.campaign.areaoptions.get
//
// 取得推广计划的可设置投放地域列表
func Taobaosimbacampaignareaoptionsget(clt *core.SDKClient, req *simba.TaobaosimbacampaignareaoptionsgetAPIRequest, session string) (*simba.TaobaosimbacampaignareaoptionsgetAPIResponse, error) {
	var resp simba.TaobaosimbacampaignareaoptionsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
