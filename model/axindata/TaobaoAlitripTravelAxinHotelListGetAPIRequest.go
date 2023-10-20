package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelListGetAPIRequest 标准酒店信息查询-阿信 API请求
// taobao.alitrip.travel.axin.hotel.list.get
//
// 阿信酒店分销基础数据查询
type TaobaoAlitripTravelAxinHotelListGetAPIRequest struct {
	model.Params
	// 城市编码
	_cityCode int64
	// 分销商id(由阿信分配)
	_distributorTid int64
	// 页大小
	_pageSize int64
	// 页码
	_pageNo int64
}

// NewTaobaoAlitripTravelAxinHotelListGetRequest 初始化TaobaoAlitripTravelAxinHotelListGetAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelListGetRequest() *TaobaoAlitripTravelAxinHotelListGetAPIRequest {
	return &TaobaoAlitripTravelAxinHotelListGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelListGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelAxinHotelListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCityCode is CityCode Setter
// 城市编码
func (r *TaobaoAlitripTravelAxinHotelListGetAPIRequest) SetCityCode(_cityCode int64) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaoAlitripTravelAxinHotelListGetAPIRequest) GetCityCode() int64 {
	return r._cityCode
}

// SetDistributorTid is DistributorTid Setter
// 分销商id(由阿信分配)
func (r *TaobaoAlitripTravelAxinHotelListGetAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelListGetAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetPageSize is PageSize Setter
// 页大小
func (r *TaobaoAlitripTravelAxinHotelListGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoAlitripTravelAxinHotelListGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaoAlitripTravelAxinHotelListGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoAlitripTravelAxinHotelListGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
