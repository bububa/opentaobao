package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordRankPriceGet 外贸直通车关键词前五名排价
// alibaba.scbp.ad.keyword.rank.price.get
//
// 外贸直通车关键词前五名排价
func AlibabaScbpAdKeywordRankPriceGet(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordRankPriceGetAPIRequest, resp *scbp.AlibabaScbpAdKeywordRankPriceGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
