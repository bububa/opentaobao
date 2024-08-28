package legalcase

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalCaseStandpointQueryref 查询推送口径信息
// alibaba.legal.case.standpoint.queryref
//
// 查询推送口径信息
func AlibabaLegalCaseStandpointQueryref(ctx context.Context, clt *core.SDKClient, req *legalcase.AlibabaLegalCaseStandpointQueryrefAPIRequest, resp *legalcase.AlibabaLegalCaseStandpointQueryrefAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
