package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytcodetobill 通过追溯码查单据
// alibaba.alihealth.drug.kyt.codetobill
//
// 通过追溯码查单据
func Alibabaalihealthdrugkytcodetobill(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytcodetobillAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytcodetobillAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytcodetobillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
