package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadtargettagmergecampaigntargettag 标签增删改
// alibaba.scbp.ad.target.tag.merge.campaign.target.tag
//
// 标签增删改
func Alibabascbpadtargettagmergecampaigntargettag(clt *core.SDKClient, req *scbp.AlibabascbpadtargettagmergecampaigntargettagAPIRequest, session string) (*scbp.AlibabascbpadtargettagmergecampaigntargettagAPIResponse, error) {
	var resp scbp.AlibabascbpadtargettagmergecampaigntargettagAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
