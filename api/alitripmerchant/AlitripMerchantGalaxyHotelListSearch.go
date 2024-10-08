package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyHotelListSearch 星河-酒店列表页搜索
// alitrip.merchant.galaxy.hotel.list.search
//
// 星河产品=酒店列表页搜索
func AlitripMerchantGalaxyHotelListSearch(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyHotelListSearchAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyHotelListSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
