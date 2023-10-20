package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyhoteldetailsearch 星河-酒店详细信息搜索
// alitrip.merchant.galaxy.hotel.detail.search
//
// 星河服务=获取雅高酒店详细信息
func Alitripmerchantgalaxyhoteldetailsearch(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyhoteldetailsearchAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyhoteldetailsearchAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyhoteldetailsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
