package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemcampaigndaybudget 获取当日投放日预算总额
// taobao.feedflow.item.campaign.daybudget
//
// 获取当日投放日预算总额
func Taobaofeedflowitemcampaigndaybudget(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemcampaigndaybudgetAPIRequest, session string) (*feedflow.TaobaofeedflowitemcampaigndaybudgetAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemcampaigndaybudgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
