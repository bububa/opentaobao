package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadtargettagestimateuv 标签人群预估
// alibaba.scbp.ad.target.tag.estimate.uv
//
// 标签人群预估
func Alibabascbpadtargettagestimateuv(clt *core.SDKClient, req *scbp.AlibabascbpadtargettagestimateuvAPIRequest, session string) (*scbp.AlibabascbpadtargettagestimateuvAPIResponse, error) {
	var resp scbp.AlibabascbpadtargettagestimateuvAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
