package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadcampaigndeleteforbiddenkeyword 删除屏蔽词
// alibaba.scbp.ad.campaign.delete.forbidden.keyword
//
// 删除屏蔽词
func Alibabascbpadcampaigndeleteforbiddenkeyword(clt *core.SDKClient, req *scbp.AlibabascbpadcampaigndeleteforbiddenkeywordAPIRequest, session string) (*scbp.AlibabascbpadcampaigndeleteforbiddenkeywordAPIResponse, error) {
	var resp scbp.AlibabascbpadcampaigndeleteforbiddenkeywordAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
