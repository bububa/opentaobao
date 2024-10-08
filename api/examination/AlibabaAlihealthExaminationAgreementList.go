package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationAgreementList isv协议获取
// alibaba.alihealth.examination.agreement.list
//
// isv协议获取
func AlibabaAlihealthExaminationAgreementList(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationAgreementListAPIRequest, resp *examination.AlibabaAlihealthExaminationAgreementListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
