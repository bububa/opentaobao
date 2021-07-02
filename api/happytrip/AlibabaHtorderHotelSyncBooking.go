package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHtorderHotelSyncBooking 未来酒店亲橙客栈预订信息同步
// alibaba.htorder.hotel.sync.booking
//
// 未来酒店亲橙客栈预订信息同步
func AlibabaHtorderHotelSyncBooking(clt *core.SDKClient, req *happytrip.AlibabaHtorderHotelSyncBookingAPIRequest, session string) (*happytrip.AlibabaHtorderHotelSyncBookingAPIResponse, error) {
	var resp happytrip.AlibabaHtorderHotelSyncBookingAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
