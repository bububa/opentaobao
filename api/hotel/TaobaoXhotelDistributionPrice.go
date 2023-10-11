package hotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotel"
)

// TaobaoXhotelDistributionPrice 飞猪分销通用酒店报价接口
// taobao.xhotel.distribution.price
//
// 飞猪分销通用酒店报价接口
func TaobaoXhotelDistributionPrice(clt *core.SDKClient, req *hotel.TaobaoXhotelDistributionPriceAPIRequest, session string) (*hotel.TaobaoXhotelDistributionPriceAPIResponse, error) {
	var resp hotel.TaobaoXhotelDistributionPriceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
