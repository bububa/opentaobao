package xhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotel"
)

// TaobaoXhotelDataServiceHotelServiceindex 酒店服务指数
// taobao.xhotel.data.service.hotel.serviceindex
//
// 酒店服务指数
func TaobaoXhotelDataServiceHotelServiceindex(clt *core.SDKClient, req *xhotel.TaobaoXhotelDataServiceHotelServiceindexAPIRequest, resp *xhotel.TaobaoXhotelDataServiceHotelServiceindexAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
