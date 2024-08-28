package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmPointConsumepoint 积分抵现
// alibaba.alsc.crm.point.consumepoint
//
// 积分抵现
func AlibabaAlscCrmPointConsumepoint(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmPointConsumepointAPIRequest, resp *alsc.AlibabaAlscCrmPointConsumepointAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
