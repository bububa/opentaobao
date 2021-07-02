package brandhub

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/brandhub"
)

// TaobaoBrandhubSpecialshowRptCampaignGet 品牌号品牌特秀计划报表数据查询
// taobao.brandhub.specialshow.rpt.campaign.get
//
// 获取品牌号品牌特秀广告campaign分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
func TaobaoBrandhubSpecialshowRptCampaignGet(clt *core.SDKClient, req *brandhub.TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest, session string) (*brandhub.TaobaoBrandhubSpecialshowRptCampaignGetAPIResponse, error) {
	var resp brandhub.TaobaoBrandhubSpecialshowRptCampaignGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
