package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupsbycampaignidGet 批量得到推广计划下的推广单元
// taobao.simba.adgroupsbycampaignid.get
//
// 根据推广计划ID分页获取推广计划下的推广单元信息
func TaobaoSimbaAdgroupsbycampaignidGet(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupsbycampaignidGetAPIRequest, resp *simba.TaobaoSimbaAdgroupsbycampaignidGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
