package legalcase

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalCaseStandpointSavestandpoint 新增反馈口径
// alibaba.legal.case.standpoint.savestandpoint
//
// 新增反馈口径 ,从外部接受反馈的口径
func AlibabaLegalCaseStandpointSavestandpoint(ctx context.Context, clt *core.SDKClient, req *legalcase.AlibabaLegalCaseStandpointSavestandpointAPIRequest, resp *legalcase.AlibabaLegalCaseStandpointSavestandpointAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
