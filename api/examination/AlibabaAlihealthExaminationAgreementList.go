package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationagreementlist isv协议获取
// alibaba.alihealth.examination.agreement.list
//
// isv协议获取
func Alibabaalihealthexaminationagreementlist(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationagreementlistAPIRequest, session string) (*examination.AlibabaalihealthexaminationagreementlistAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationagreementlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
