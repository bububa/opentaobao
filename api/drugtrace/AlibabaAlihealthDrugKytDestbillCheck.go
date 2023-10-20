package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdestbillcheck 直调审批
// alibaba.alihealth.drug.kyt.destbill.check
//
// 为药企提供直调单据审批操作
func Alibabaalihealthdrugkytdestbillcheck(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdestbillcheckAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdestbillcheckAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdestbillcheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
