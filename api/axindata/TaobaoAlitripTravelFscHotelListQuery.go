package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscHotelListQuery 标准酒店信息查询-供销平台
// taobao.alitrip.travel.fsc.hotel.list.query
//
// 供销平台标准酒店信息列表查询
func TaobaoAlitripTravelFscHotelListQuery(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscHotelListQueryAPIRequest, resp *axindata.TaobaoAlitripTravelFscHotelListQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
