package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelHstdfShotelExnotmatchroomAPIRequest
导出一个hid下所有未匹配rid的接口 API请求
alitrip.hotel.hstdf.shotel.exnotmatchroom

导出一个卖家hid下所有未匹配的rid信息，包括rid，名称、英文名称、床型、窗型、面积、对外系统id */
type AlitripHotelHstdfShotelExnotmatchroomAPIRequest struct {
	model.Params
	// 卖家酒店hid
	_hid int64
}

// New
