package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyHotelDetailSearchData 星河-酒店详情页信息获取(新改版)
// alitrip.merchant.galaxy.hotel.detail.search.data
//
// 星河服务，获取雅高酒店详细信息，详情页新改版接口
func AlitripMerchantGalaxyHotelDetailSearchData(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyHotelDetailSearchDataAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyHotelDetailSearchDataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
