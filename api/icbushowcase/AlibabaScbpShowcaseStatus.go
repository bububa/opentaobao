package icbushowcase

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

// AlibabaScbpShowcaseStatus 橱窗状态
// alibaba.scbp.showcase.status
//
// 查询橱窗状态，如总数、可用数量
func AlibabaScbpShowcaseStatus(ctx context.Context, clt *core.SDKClient, req *icbushowcase.AlibabaScbpShowcaseStatusAPIRequest, resp *icbushowcase.AlibabaScbpShowcaseStatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
