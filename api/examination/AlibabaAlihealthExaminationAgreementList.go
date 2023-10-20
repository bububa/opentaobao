package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationAgreementList isv协议获取
// alibaba.alihealth.examination.agreement.list
//
// isv协议获取
func AlibabaAlihealthExaminationAgreementList(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationAgreementListAPIRequest, resp *examination.AlibabaAlihealthExaminationAgreementListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
