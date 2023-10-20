package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyHotelListSearch 星河-酒店列表页搜索
// alitrip.merchant.galaxy.hotel.list.search
//
// 星河产品=酒店列表页搜索
func AlitripMerchantGalaxyHotelListSearch(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyHotelListSearchAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyHotelListSearchAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyHotelListSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
