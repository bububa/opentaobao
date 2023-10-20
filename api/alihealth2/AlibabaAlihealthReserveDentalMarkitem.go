package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthreservedentalmarkitem 标记商品是否可预约
// alibaba.alihealth.reserve.dental.markitem
//
// 标记商品是否可预约
func Alibabaalihealthreservedentalmarkitem(clt *core.SDKClient, req *alihealth2.AlibabaalihealthreservedentalmarkitemAPIRequest, session string) (*alihealth2.AlibabaalihealthreservedentalmarkitemAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthreservedentalmarkitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
