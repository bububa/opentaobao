package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadtargettaglistrecommendtag 给计划推荐标签
// alibaba.scbp.ad.target.tag.list.recommend.tag
//
// 给计划推荐标签
func Alibabascbpadtargettaglistrecommendtag(clt *core.SDKClient, req *scbp.AlibabascbpadtargettaglistrecommendtagAPIRequest, session string) (*scbp.AlibabascbpadtargettaglistrecommendtagAPIResponse, error) {
	var resp scbp.AlibabascbpadtargettaglistrecommendtagAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
