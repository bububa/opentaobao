package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyorderfill 填单页接口
// alitrip.merchant.galaxy.order.fill
//
// 进入填单页时调用此接口，返回填单所需展示基础信息
func Alitripmerchantgalaxyorderfill(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyorderfillAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyorderfillAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyorderfillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
