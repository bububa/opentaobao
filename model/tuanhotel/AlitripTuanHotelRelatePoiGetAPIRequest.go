package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTuanHotelRelatePoiGetAPIRequest 酒店团购关联POI API请求
// alitrip.tuan.hotel.relate.poi.get
//
// 查询酒店团购关联POI
type AlitripTuanHotelRelatePoiGetAPIRequest struct {
	model.Params
	// 关键字
	_keywords string
	// 行政区ID
	_divisionId int64
	// 类目ID：国内酒店套餐-201189402；国际酒店套餐-201188002；酒店餐饮美食-201214101；酒店服务-201214201；酒店客房优惠券-201214301
	_catId int64
}

// NewAlitripTuanHotelRelatePoiGetRequest 初始化AlitripTuanHotelRelatePoiGetAPIRequest对象
func NewAlitripTuanHotelRelatePoiGetRequest() *AlitripTuanHotelRelatePoiGetAPIRequest {
	return &AlitripTuanHotelRelatePoiGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTuanHotelRelatePoiGetAPIRequest) GetApiMethodName() string {
	return "alitrip.tuan.hotel.relate.poi.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTuanHotelRelatePoiGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Keywords Setter
// 关键字
func (r *AlitripTuanHotelRelatePoiGetAPIRequest) SetKeywords(_keywords string) error {
	r._keywords = _keywords
	r.Set("keywords", _keywords)
	return nil
}

// Get Keywords Getter
func (r AlitripTuanHotelRelatePoiGetAPIRequest) GetKeywords() string {
	return r._keywords
}

// Set is DivisionId Setter
// 行政区ID
func (r *AlitripTuanHotelRelatePoiGetAPIRequest) SetDivisionId(_divisionId int64) error {
	r._divisionId = _divisionId
	r.Set("division_id", _divisionId)
	return nil
}

// Get DivisionId Getter
func (r AlitripTuanHotelRelatePoiGetAPIRequest) GetDivisionId() int64 {
	return r._divisionId
}

// Set is CatId Setter
// 类目ID：国内酒店套餐-201189402；国际酒店套餐-201188002；酒店餐饮美食-201214101；酒店服务-201214201；酒店客房优惠券-201214301
func (r *AlitripTuanHotelRelatePoiGetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// Get CatId Getter
func (r AlitripTuanHotelRelatePoiGetAPIRequest) GetCatId() int64 {
	return r._catId
}
