package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadcampaigndelete 删除计划
// alibaba.scbp.ad.campaign.delete
//
// 删除计划
func Alibabascbpadcampaigndelete(clt *core.SDKClient, req *scbp.AlibabascbpadcampaigndeleteAPIRequest, session string) (*scbp.AlibabascbpadcampaigndeleteAPIResponse, error) {
	var resp scbp.AlibabascbpadcampaigndeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
