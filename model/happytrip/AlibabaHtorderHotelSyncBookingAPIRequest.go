package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHtorderHotelSyncBookingAPIRequest
未来酒店亲橙客栈预订信息同步 API请求
alibaba.htorder.hotel.sync.booking

未来酒店亲橙客栈预订信息同步 */
type AlibabaHtorderHotelSyncBookingAPIRequest struct {
	model.Params
	// 预订信息数据
	_dataEntity *SyncHotelBookingDataRequestDto
}

// New
