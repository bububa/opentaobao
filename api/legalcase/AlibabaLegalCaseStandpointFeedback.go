package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

/* AlibabaLegalCaseStandpointFeedback
新增或更新 反馈口径(采纳口径/不采纳口径)
alibaba.legal.case.standpoint.feedback

新增或更新 反馈口径(采纳口径/不采纳口径) */
func AlibabaLegalCaseStandpointFeedback(clt *core.SDKClient, req *legalcase.AlibabaLegalCaseStandpointFeedbackAPIRequest, session string) (*legalcase.AlibabaLegalCaseStandpointFeedbackAPIResponse, error) {
	var resp legalcase.AlibabaLegalCaseStandpointFeedbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
