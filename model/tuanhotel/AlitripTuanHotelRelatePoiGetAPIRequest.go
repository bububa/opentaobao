package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTuanHotelRelatePoiGetAPIRequest
酒店团购关联POI API请求
alitrip.tuan.hotel.relate.poi.get

查询酒店团购关联POI */
type AlitripTuanHotelRelatePoiGetAPIRequest struct {
	model.Params
	// 关键字
	_keywords string
	// 行政区ID
	_divisionId int64
	// 类目ID：国内酒店套餐-201189402；国际酒店套餐-201188002；酒店餐饮美食-201214101；酒店服务-201214201；酒店客房优惠券-201214301
	_catId int64
}

// New
