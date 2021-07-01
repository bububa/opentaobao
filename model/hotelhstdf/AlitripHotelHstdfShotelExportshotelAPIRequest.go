package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelHstdfShotelExportshotelAPIRequest
商家自主导出相似度高的标准酒店 API请求
alitrip.hotel.hstdf.shotel.exportshotel

商家通过给出自己的卖家酒店信息，通过接口可以返回相似度高的标准酒店信息 */
type AlitripHotelHstdfShotelExportshotelAPIRequest struct {
	model.Params
	// HID，卖家酒店上传到平台后的ID
	_hid int64
	// 酒店名称，必填
	_name string
	// 酒店所在行政区划，对应平台ID，为空时会使用经纬度来定位
	_cityCode int64
	// 电话
	_telNumber string
	// 经度
	_longitude string
	// 纬度
	_latitude string
	// 地址
	_address string
}

// New
