package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripHotelOrderSearch 搜索酒店订单列表
// alitrip.btrip.hotel.order.search
//
// 企业获取商旅酒店订单数据
func AlitripBtripHotelOrderSearch(clt *core.SDKClient, req *btrip.AlitripBtripHotelOrderSearchAPIRequest, session string) (*btrip.AlitripBtripHotelOrderSearchAPIResponse, error) {
	var resp btrip.AlitripBtripHotelOrderSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
