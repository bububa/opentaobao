package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytyydrugcodes 查询药品是否赋码
// alibaba.alihealth.drug.kyt.yy.drugcodes
//
// 药品是否赋码
func Alibabaalihealthdrugkytyydrugcodes(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytyydrugcodesAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytyydrugcodesAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytyydrugcodesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
