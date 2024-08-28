package legalcase

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalCaseEntrustCallback 委托回调接口
// alibaba.legal.case.entrust.callback
//
// 委托回调接口
func AlibabaLegalCaseEntrustCallback(ctx context.Context, clt *core.SDKClient, req *legalcase.AlibabaLegalCaseEntrustCallbackAPIRequest, resp *legalcase.AlibabaLegalCaseEntrustCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
