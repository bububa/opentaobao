package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordRankPriceGet 外贸直通车关键词前五名排价
// alibaba.scbp.ad.keyword.rank.price.get
//
// 外贸直通车关键词前五名排价
func AlibabaScbpAdKeywordRankPriceGet(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordRankPriceGetAPIRequest, resp *scbp.AlibabaScbpAdKeywordRankPriceGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
