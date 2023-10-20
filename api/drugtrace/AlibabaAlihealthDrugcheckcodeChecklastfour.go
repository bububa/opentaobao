package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcheckcodechecklastfour 校验追溯码的后4位是否正确
// alibaba.alihealth.drugcheckcode.checklastfour
//
// 校验追溯码的后4位是否正确
func Alibabaalihealthdrugcheckcodechecklastfour(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcheckcodechecklastfourAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcheckcodechecklastfourAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcheckcodechecklastfourAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
