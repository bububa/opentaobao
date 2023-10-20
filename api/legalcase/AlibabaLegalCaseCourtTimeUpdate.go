package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalCaseCourtTimeUpdate 开庭时间变更
// alibaba.legal.case.court.time.update
//
// 修改案件的开庭时间
func AlibabaLegalCaseCourtTimeUpdate(clt *core.SDKClient, req *legalcase.AlibabaLegalCaseCourtTimeUpdateAPIRequest, resp *legalcase.AlibabaLegalCaseCourtTimeUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
