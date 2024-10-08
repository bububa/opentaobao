package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordRankPriceBatchget 外贸直通车关键词前五名批量排价
// alibaba.scbp.ad.keyword.rank.price.batchget
//
// 外贸直通车关键词前五名批量排价
func AlibabaScbpAdKeywordRankPriceBatchget(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordRankPriceBatchgetAPIRequest, resp *scbp.AlibabaScbpAdKeywordRankPriceBatchgetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
