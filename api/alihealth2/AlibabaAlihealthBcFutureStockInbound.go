package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthbcfuturestockinbound 供应商上报期货库存
// alibaba.alihealth.bc.future.stock.inbound
//
// 供应商上报期货库存
func Alibabaalihealthbcfuturestockinbound(clt *core.SDKClient, req *alihealth2.AlibabaalihealthbcfuturestockinboundAPIRequest, session string) (*alihealth2.AlibabaalihealthbcfuturestockinboundAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthbcfuturestockinboundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
