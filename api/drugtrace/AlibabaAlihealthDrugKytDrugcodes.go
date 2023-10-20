package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdrugcodes 药品是否赋码
// alibaba.alihealth.drug.kyt.drugcodes
//
// 药品是否赋码
func Alibabaalihealthdrugkytdrugcodes(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdrugcodesAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdrugcodesAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdrugcodesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
