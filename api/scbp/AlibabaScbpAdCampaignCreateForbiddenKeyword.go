package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadcampaigncreateforbiddenkeyword 创建屏蔽词
// alibaba.scbp.ad.campaign.create.forbidden.keyword
//
// 创建屏蔽词
func Alibabascbpadcampaigncreateforbiddenkeyword(clt *core.SDKClient, req *scbp.AlibabascbpadcampaigncreateforbiddenkeywordAPIRequest, session string) (*scbp.AlibabascbpadcampaigncreateforbiddenkeywordAPIResponse, error) {
	var resp scbp.AlibabascbpadcampaigncreateforbiddenkeywordAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
