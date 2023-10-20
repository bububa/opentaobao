package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordRankGet 获取外贸直通车关键词预估排名
// alibaba.scbp.ad.keyword.rank.get
//
// 获取外贸直通车关键词预估排名
func AlibabaScbpAdKeywordRankGet(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordRankGetAPIRequest, resp *scbp.AlibabaScbpAdKeywordRankGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
