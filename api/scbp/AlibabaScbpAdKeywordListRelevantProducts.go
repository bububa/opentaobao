package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordListRelevantProducts 查询和词匹配的推广产品
// alibaba.scbp.ad.keyword.list.relevant.products
//
// 查询和词匹配的推广产品
func AlibabaScbpAdKeywordListRelevantProducts(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordListRelevantProductsAPIRequest, resp *scbp.AlibabaScbpAdKeywordListRelevantProductsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
