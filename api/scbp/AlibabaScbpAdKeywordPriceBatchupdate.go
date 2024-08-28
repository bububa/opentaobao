package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordPriceBatchupdate 关键词批量改价
// alibaba.scbp.ad.keyword.price.batchupdate
//
// 关键词批量改价
func AlibabaScbpAdKeywordPriceBatchupdate(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordPriceBatchupdateAPIRequest, resp *scbp.AlibabaScbpAdKeywordPriceBatchupdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
