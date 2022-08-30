package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordRecommendWord 推词
// alibaba.scbp.ad.keyword.recommend.word
//
// 推词
func AlibabaScbpAdKeywordRecommendWord(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordRecommendWordAPIRequest, session string) (*scbp.AlibabaScbpAdKeywordRecommendWordAPIResponse, error) {
	var resp scbp.AlibabaScbpAdKeywordRecommendWordAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
