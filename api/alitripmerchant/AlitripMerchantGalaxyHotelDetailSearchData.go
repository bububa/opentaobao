package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyhoteldetailsearchdata 星河-酒店详情页信息获取(新改版)
// alitrip.merchant.galaxy.hotel.detail.search.data
//
// 星河服务，获取雅高酒店详细信息，详情页新改版接口
func Alitripmerchantgalaxyhoteldetailsearchdata(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyhoteldetailsearchdataAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyhoteldetailsearchdataAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyhoteldetailsearchdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
