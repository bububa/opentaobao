package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordUpdateKeywordPriceBatch 修改关键词价格
// alibaba.scbp.ad.keyword.update.keyword.price.batch
//
// 修改关键词价格
func AlibabaScbpAdKeywordUpdateKeywordPriceBatch(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest, resp *scbp.AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
