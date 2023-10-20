package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadcampaignfindrealcost 批量查询计划消耗数据
// alibaba.scbp.ad.campaign.find.real.cost
//
// 批量查询计划消耗数据
func Alibabascbpadcampaignfindrealcost(clt *core.SDKClient, req *scbp.AlibabascbpadcampaignfindrealcostAPIRequest, session string) (*scbp.AlibabascbpadcampaignfindrealcostAPIResponse, error) {
	var resp scbp.AlibabascbpadcampaignfindrealcostAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
