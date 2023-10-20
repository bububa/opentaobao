package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyordervalidate 星河-订单试单接口
// alitrip.merchant.galaxy.order.validate
//
// 根据用户选择酒店房型、入住人数、预订时间参数，获取是否可预订及价格变化信息
func Alitripmerchantgalaxyordervalidate(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyordervalidateAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyordervalidateAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyordervalidateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
