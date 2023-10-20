package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalCaseEntrustGet 委托
// alibaba.legal.case.entrust.get
//
// 获取委托案件的基本信息
func AlibabaLegalCaseEntrustGet(clt *core.SDKClient, req *legalcase.AlibabaLegalCaseEntrustGetAPIRequest, resp *legalcase.AlibabaLegalCaseEntrustGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
