package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtriphotelordersearch 搜索酒店订单列表
// alitrip.btrip.hotel.order.search
//
// 企业获取商旅酒店订单数据
func Alitripbtriphotelordersearch(clt *core.SDKClient, req *btrip.AlitripbtriphotelordersearchAPIRequest, session string) (*btrip.AlitripbtriphotelordersearchAPIResponse, error) {
	var resp btrip.AlitripbtriphotelordersearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
