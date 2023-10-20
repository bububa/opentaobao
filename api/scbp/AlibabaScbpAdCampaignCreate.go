package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadcampaigncreate 创建计划
// alibaba.scbp.ad.campaign.create
//
// 创建计划
func Alibabascbpadcampaigncreate(clt *core.SDKClient, req *scbp.AlibabascbpadcampaigncreateAPIRequest, session string) (*scbp.AlibabascbpadcampaigncreateAPIResponse, error) {
	var resp scbp.AlibabascbpadcampaigncreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
