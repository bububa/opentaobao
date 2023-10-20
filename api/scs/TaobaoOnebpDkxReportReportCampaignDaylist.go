package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// Taobaoonebpdkxreportreportcampaigndaylist 获取计划分日报表
// taobao.onebp.dkx.report.report.campaign.daylist
//
// 获取计划分日报表，场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe
func Taobaoonebpdkxreportreportcampaigndaylist(clt *core.SDKClient, req *scs.TaobaoonebpdkxreportreportcampaigndaylistAPIRequest, session string) (*scs.TaobaoonebpdkxreportreportcampaigndaylistAPIResponse, error) {
	var resp scs.TaobaoonebpdkxreportreportcampaigndaylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
