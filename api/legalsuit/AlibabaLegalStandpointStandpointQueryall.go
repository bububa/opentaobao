package legalsuit

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointStandpointQueryall 滑动查询口径
// alibaba.legal.standpoint.standpoint.queryall
//
// 滑动查询口径
func AlibabaLegalStandpointStandpointQueryall(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointStandpointQueryallAPIRequest, resp *legalsuit.AlibabaLegalStandpointStandpointQueryallAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
