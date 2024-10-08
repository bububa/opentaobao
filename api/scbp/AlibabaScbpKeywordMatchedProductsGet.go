package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpKeywordMatchedProductsGet 查询和词匹配的推广产品
// alibaba.scbp.keyword.matched.products.get
//
// 查询和词匹配的推广产品
func AlibabaScbpKeywordMatchedProductsGet(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpKeywordMatchedProductsGetAPIRequest, resp *scbp.AlibabaScbpKeywordMatchedProductsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
