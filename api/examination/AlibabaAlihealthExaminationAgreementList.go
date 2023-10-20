package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationAgreementList isv协议获取
// alibaba.alihealth.examination.agreement.list
//
// isv协议获取
func AlibabaAlihealthExaminationAgreementList(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationAgreementListAPIRequest, session string) (*examination.AlibabaAlihealthExaminationAgreementListAPIResponse, error) {
	var resp examination.AlibabaAlihealthExaminationAgreementListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
