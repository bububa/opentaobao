package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthbcitemperiodsync 代销品效期同步
// alibaba.alihealth.bc.item.period.sync
//
// 代销品效期同步
func Alibabaalihealthbcitemperiodsync(clt *core.SDKClient, req *alihealth2.AlibabaalihealthbcitemperiodsyncAPIRequest, session string) (*alihealth2.AlibabaalihealthbcitemperiodsyncAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthbcitemperiodsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
