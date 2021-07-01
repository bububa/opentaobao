package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTuanHotelItemInfoGetAPIRequest
宝贝信息查询接口 API请求
alitrip.tuan.hotel.item.info.get

商家查询发布的宝贝详情信息 */
type AlitripTuanHotelItemInfoGetAPIRequest struct {
	model.Params
	// 宝贝ID
	_itemId int64
}

// New
