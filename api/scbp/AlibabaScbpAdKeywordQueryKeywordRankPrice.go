package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordQueryKeywordRankPrice 查询关键词前五名排价
// alibaba.scbp.ad.keyword.query.keyword.rank.price
//
// 查询关键词前五名排价
func AlibabaScbpAdKeywordQueryKeywordRankPrice(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest, resp *scbp.AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
