package legalsuit

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointStandpointtreeQuery 查询口径树目录
// alibaba.legal.standpoint.standpointtree.query
//
// 查询口径树目录
func AlibabaLegalStandpointStandpointtreeQuery(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointStandpointtreeQueryAPIRequest, resp *legalsuit.AlibabaLegalStandpointStandpointtreeQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
