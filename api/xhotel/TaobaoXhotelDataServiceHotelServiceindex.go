package xhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotel"
)

/* TaobaoXhotelDataServiceHotelServiceindex
酒店服务指数
taobao.xhotel.data.service.hotel.serviceindex

酒店服务指数 */
func TaobaoXhotelDataServiceHotelServiceindex(clt *core.SDKClient, req *xhotel.TaobaoXhotelDataServiceHotelServiceindexAPIRequest, session string) (*xhotel.TaobaoXhotelDataServiceHotelServiceindexAPIResponse, error) {
	var resp xhotel.TaobaoXhotelDataServiceHotelServiceindexAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
