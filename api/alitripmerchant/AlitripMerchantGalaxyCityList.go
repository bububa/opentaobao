package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyCityList 星河-酒店城市列表展示
// alitrip.merchant.galaxy.city.list
//
// 雅高酒店城市列表展示，并且首字母列出酒店城市
func AlitripMerchantGalaxyCityList(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyCityListAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyCityListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
