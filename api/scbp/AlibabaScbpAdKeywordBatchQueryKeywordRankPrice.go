package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordBatchQueryKeywordRankPrice 批量查询关键词前五名排价
// alibaba.scbp.ad.keyword.batch.query.keyword.rank.price
//
// 批量查询关键词前五名排价
func AlibabaScbpAdKeywordBatchQueryKeywordRankPrice(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest, resp *scbp.AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
