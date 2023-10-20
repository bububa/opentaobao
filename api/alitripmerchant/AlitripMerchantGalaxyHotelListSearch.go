package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyhotellistsearch 星河-酒店列表页搜索
// alitrip.merchant.galaxy.hotel.list.search
//
// 星河产品=酒店列表页搜索
func Alitripmerchantgalaxyhotellistsearch(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyhotellistsearchAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyhotellistsearchAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyhotellistsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
