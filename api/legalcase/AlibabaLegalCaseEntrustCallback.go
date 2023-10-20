package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalCaseEntrustCallback 委托回调接口
// alibaba.legal.case.entrust.callback
//
// 委托回调接口
func AlibabaLegalCaseEntrustCallback(clt *core.SDKClient, req *legalcase.AlibabaLegalCaseEntrustCallbackAPIRequest, resp *legalcase.AlibabaLegalCaseEntrustCallbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
