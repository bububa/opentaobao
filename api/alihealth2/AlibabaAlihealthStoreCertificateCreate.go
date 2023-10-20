package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthstorecertificatecreate 仓库换证审批
// alibaba.alihealth.store.certificate.create
//
// 仓库侧换证发起审批
func Alibabaalihealthstorecertificatecreate(clt *core.SDKClient, req *alihealth2.AlibabaalihealthstorecertificatecreateAPIRequest, session string) (*alihealth2.AlibabaalihealthstorecertificatecreateAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthstorecertificatecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
