package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyOrderFill 填单页接口
// alitrip.merchant.galaxy.order.fill
//
// 进入填单页时调用此接口，返回填单所需展示基础信息
func AlitripMerchantGalaxyOrderFill(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyOrderFillAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyOrderFillAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyOrderFillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
