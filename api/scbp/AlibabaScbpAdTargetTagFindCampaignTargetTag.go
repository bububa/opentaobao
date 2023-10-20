package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadtargettagfindcampaigntargettag 查询标签数据
// alibaba.scbp.ad.target.tag.find.campaign.target.tag
//
// 查询标签数据
func Alibabascbpadtargettagfindcampaigntargettag(clt *core.SDKClient, req *scbp.AlibabascbpadtargettagfindcampaigntargettagAPIRequest, session string) (*scbp.AlibabascbpadtargettagfindcampaigntargettagAPIResponse, error) {
	var resp scbp.AlibabascbpadtargettagfindcampaigntargettagAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
