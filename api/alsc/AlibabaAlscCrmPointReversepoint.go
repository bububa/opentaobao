package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmPointReversepoint 积分消费回退
// alibaba.alsc.crm.point.reversepoint
//
// 积分消费回退
func AlibabaAlscCrmPointReversepoint(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmPointReversepointAPIRequest, resp *alsc.AlibabaAlscCrmPointReversepointAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
