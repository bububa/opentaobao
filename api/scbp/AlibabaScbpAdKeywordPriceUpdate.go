package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordPriceUpdate 关键词改价
// alibaba.scbp.ad.keyword.price.update
//
// 关键词改价
func AlibabaScbpAdKeywordPriceUpdate(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordPriceUpdateAPIRequest, resp *scbp.AlibabaScbpAdKeywordPriceUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
