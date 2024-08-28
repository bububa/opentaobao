package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyHotelDetailSearch 星河-酒店详细信息搜索
// alitrip.merchant.galaxy.hotel.detail.search
//
// 星河服务=获取雅高酒店详细信息
func AlitripMerchantGalaxyHotelDetailSearch(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyHotelDetailSearchAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyHotelDetailSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
