package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthbcfuturestockoutbound 供应商期货出库
// alibaba.alihealth.bc.future.stock.outbound
//
// 供应商期货出库
func Alibabaalihealthbcfuturestockoutbound(clt *core.SDKClient, req *alihealth2.AlibabaalihealthbcfuturestockoutboundAPIRequest, session string) (*alihealth2.AlibabaalihealthbcfuturestockoutboundAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthbcfuturestockoutboundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
