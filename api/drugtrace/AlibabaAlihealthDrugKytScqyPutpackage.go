package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytscqyputpackage 码拼箱
// alibaba.alihealth.drug.kyt.scqy.putpackage
//
// 码拼箱接口
func Alibabaalihealthdrugkytscqyputpackage(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytscqyputpackageAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytscqyputpackageAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytscqyputpackageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
