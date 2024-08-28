package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpProductList 查询P4P产品
// alibaba.scbp.product.list
//
// 查询P4P产品
func AlibabaScbpProductList(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpProductListAPIRequest, resp *scbp.AlibabaScbpProductListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
