package hotelalliance

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelAllianceHidGetAPIRequest
获取联盟hid API请求
alitrip.hotel.alliance.hid.get

获取符合条件的菲住联盟hid，目前支持指定日期上线的菲住联盟hid查询 */
type AlitripHotelAllianceHidGetAPIRequest struct {
	model.Params
	// 查询入参
	_allianceInfoRequest *AllianceInfoRequest
}

// New
