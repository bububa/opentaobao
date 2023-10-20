package hotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotel"
)

// Taobaoxhoteldistributionprice 飞猪分销通用酒店报价接口
// taobao.xhotel.distribution.price
//
// 飞猪分销通用酒店报价接口
func Taobaoxhoteldistributionprice(clt *core.SDKClient, req *hotel.TaobaoxhoteldistributionpriceAPIRequest, session string) (*hotel.TaobaoxhoteldistributionpriceAPIResponse, error) {
	var resp hotel.TaobaoxhoteldistributionpriceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
