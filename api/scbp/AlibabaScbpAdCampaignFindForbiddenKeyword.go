package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadcampaignfindforbiddenkeyword 查询屏蔽词
// alibaba.scbp.ad.campaign.find.forbidden.keyword
//
// 查询屏蔽词
func Alibabascbpadcampaignfindforbiddenkeyword(clt *core.SDKClient, req *scbp.AlibabascbpadcampaignfindforbiddenkeywordAPIRequest, session string) (*scbp.AlibabascbpadcampaignfindforbiddenkeywordAPIResponse, error) {
	var resp scbp.AlibabascbpadcampaignfindforbiddenkeywordAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
