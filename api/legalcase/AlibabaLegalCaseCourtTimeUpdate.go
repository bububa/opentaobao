package legalcase

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalCaseCourtTimeUpdate 开庭时间变更
// alibaba.legal.case.court.time.update
//
// 修改案件的开庭时间
func AlibabaLegalCaseCourtTimeUpdate(ctx context.Context, clt *core.SDKClient, req *legalcase.AlibabaLegalCaseCourtTimeUpdateAPIRequest, resp *legalcase.AlibabaLegalCaseCourtTimeUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
