package brandhub

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/brandhub"
)

// Taobaobrandstartshoprptcampaignget 明星店铺推广计划报表数据查询
// taobao.brand.startshop.rpt.campaign.get
//
// 获取明星店铺广告campaign分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
func Taobaobrandstartshoprptcampaignget(clt *core.SDKClient, req *brandhub.TaobaobrandstartshoprptcampaigngetAPIRequest, session string) (*brandhub.TaobaobrandstartshoprptcampaigngetAPIResponse, error) {
	var resp brandhub.TaobaobrandstartshoprptcampaigngetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
