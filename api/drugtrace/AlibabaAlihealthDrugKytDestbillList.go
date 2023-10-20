package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdestbilllist 直调单据查询
// alibaba.alihealth.drug.kyt.destbill.list
//
// 为药企提供直调单据查询功能
func Alibabaalihealthdrugkytdestbilllist(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdestbilllistAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdestbilllistAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdestbilllistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
