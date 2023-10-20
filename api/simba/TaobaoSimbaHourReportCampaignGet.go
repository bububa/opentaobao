package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbahourreportcampaignget 计划维度小时报表获取
// taobao.simba.hour.report.campaign.get
//
// 计划维度小时报表获取
func Taobaosimbahourreportcampaignget(clt *core.SDKClient, req *simba.TaobaosimbahourreportcampaigngetAPIRequest, session string) (*simba.TaobaosimbahourreportcampaigngetAPIResponse, error) {
	var resp simba.TaobaosimbahourreportcampaigngetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
