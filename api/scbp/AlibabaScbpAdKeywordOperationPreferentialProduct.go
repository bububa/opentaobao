package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordOperationPreferentialProduct 操作优推品
// alibaba.scbp.ad.keyword.operation.preferential.product
//
// 操作优推品
func AlibabaScbpAdKeywordOperationPreferentialProduct(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest, resp *scbp.AlibabaScbpAdKeywordOperationPreferentialProductAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
