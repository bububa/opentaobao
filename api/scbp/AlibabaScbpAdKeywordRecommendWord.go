package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordRecommendWord 推词
// alibaba.scbp.ad.keyword.recommend.word
//
// 推词
func AlibabaScbpAdKeywordRecommendWord(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordRecommendWordAPIRequest, resp *scbp.AlibabaScbpAdKeywordRecommendWordAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
