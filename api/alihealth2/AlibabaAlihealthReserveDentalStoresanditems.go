package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthreservedentalstoresanditems 查询商户门店，商品列表
// alibaba.alihealth.reserve.dental.storesanditems
//
// 查询商户门店，商品列表
func Alibabaalihealthreservedentalstoresanditems(clt *core.SDKClient, req *alihealth2.AlibabaalihealthreservedentalstoresanditemsAPIRequest, session string) (*alihealth2.AlibabaalihealthreservedentalstoresanditemsAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthreservedentalstoresanditemsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
