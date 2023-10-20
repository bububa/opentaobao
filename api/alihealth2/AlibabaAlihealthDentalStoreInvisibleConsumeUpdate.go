package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthdentalstoreinvisibleconsumeupdate 门店无隐形消费签约
// alibaba.alihealth.dental.store.invisible.consume.update
//
// 门店无隐形消费签约
func Alibabaalihealthdentalstoreinvisibleconsumeupdate(clt *core.SDKClient, req *alihealth2.AlibabaalihealthdentalstoreinvisibleconsumeupdateAPIRequest, session string) (*alihealth2.AlibabaalihealthdentalstoreinvisibleconsumeupdateAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthdentalstoreinvisibleconsumeupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
