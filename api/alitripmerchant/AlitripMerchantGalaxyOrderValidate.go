package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyOrderValidate 星河-订单试单接口
// alitrip.merchant.galaxy.order.validate
//
// 根据用户选择酒店房型、入住人数、预订时间参数，获取是否可预订及价格变化信息
func AlitripMerchantGalaxyOrderValidate(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyOrderValidateAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyOrderValidateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
