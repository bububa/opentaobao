package icbushowcase

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

// AlibabaScbpShowcaseDeleteproduct 批量删除橱窗商品
// alibaba.scbp.showcase.deleteproduct
//
// 批量删除橱窗商品
func AlibabaScbpShowcaseDeleteproduct(ctx context.Context, clt *core.SDKClient, req *icbushowcase.AlibabaScbpShowcaseDeleteproductAPIRequest, resp *icbushowcase.AlibabaScbpShowcaseDeleteproductAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
