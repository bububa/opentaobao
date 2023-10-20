package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthdentalstoreinsertorupdate ISV新增/修改口腔门店
// alibaba.alihealth.dental.store.insertorupdate
//
// ISV新增/修改口腔门店
func Alibabaalihealthdentalstoreinsertorupdate(clt *core.SDKClient, req *alihealth2.AlibabaalihealthdentalstoreinsertorupdateAPIRequest, session string) (*alihealth2.AlibabaalihealthdentalstoreinsertorupdateAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthdentalstoreinsertorupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
