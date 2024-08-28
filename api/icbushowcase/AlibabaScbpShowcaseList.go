package icbushowcase

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

// AlibabaScbpShowcaseList 橱窗查询
// alibaba.scbp.showcase.list
//
// 橱窗查询
func AlibabaScbpShowcaseList(ctx context.Context, clt *core.SDKClient, req *icbushowcase.AlibabaScbpShowcaseListAPIRequest, resp *icbushowcase.AlibabaScbpShowcaseListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
