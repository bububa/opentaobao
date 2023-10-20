package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthdentalbindauditquery ISV查询绑定审核状态
// alibaba.alihealth.dental.bind.audit.query
//
// ISV查询绑定审核状态
func Alibabaalihealthdentalbindauditquery(clt *core.SDKClient, req *alihealth2.AlibabaalihealthdentalbindauditqueryAPIRequest, session string) (*alihealth2.AlibabaalihealthdentalbindauditqueryAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthdentalbindauditqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
