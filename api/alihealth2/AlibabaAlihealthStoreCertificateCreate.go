package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthStoreCertificateCreate 仓库换证审批
// alibaba.alihealth.store.certificate.create
//
// 仓库侧换证发起审批
func AlibabaAlihealthStoreCertificateCreate(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthStoreCertificateCreateAPIRequest, resp *alihealth2.AlibabaAlihealthStoreCertificateCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
