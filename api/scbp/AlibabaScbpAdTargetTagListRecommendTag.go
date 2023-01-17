package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdTargetTagListRecommendTag 给计划推荐标签
// alibaba.scbp.ad.target.tag.list.recommend.tag
//
// 给计划推荐标签
func AlibabaScbpAdTargetTagListRecommendTag(clt *core.SDKClient, req *scbp.AlibabaScbpAdTargetTagListRecommendTagAPIRequest, session string) (*scbp.AlibabaScbpAdTargetTagListRecommendTagAPIResponse, error) {
	var resp scbp.AlibabaScbpAdTargetTagListRecommendTagAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
