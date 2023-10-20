package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadcampaignupdate 修改计划
// alibaba.scbp.ad.campaign.update
//
// 修改计划
func Alibabascbpadcampaignupdate(clt *core.SDKClient, req *scbp.AlibabascbpadcampaignupdateAPIRequest, session string) (*scbp.AlibabascbpadcampaignupdateAPIResponse, error) {
	var resp scbp.AlibabascbpadcampaignupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
