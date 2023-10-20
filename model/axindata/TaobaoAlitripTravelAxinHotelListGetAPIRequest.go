package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotellistgetAPIRequest 标准酒店信息查询-阿信 API请求
// taobao.alitrip.travel.axin.hotel.list.get
//
// 阿信酒店分销基础数据查询
type TaobaoalitriptravelaxinhotellistgetAPIRequest struct {
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

// NewTaobaoalitriptravelaxinhotellistgetRequest 初始化TaobaoalitriptravelaxinhotellistgetAPIRequest对象
func NewTaobaoalitriptravelaxinhotellistgetRequest() *TaobaoalitriptravelaxinhotellistgetAPIRequest {
	return &TaobaoalitriptravelaxinhotellistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinhotellistgetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinhotellistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinhotellistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCityCode is CityCode Setter
// 城市编码
func (r *TaobaoalitriptravelaxinhotellistgetAPIRequest) SetCityCode(_cityCode int64) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaoalitriptravelaxinhotellistgetAPIRequest) GetCityCode() int64 {
	return r._cityCode
}

// SetDistributorTid is DistributorTid Setter
// 分销商id(由阿信分配)
func (r *TaobaoalitriptravelaxinhotellistgetAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoalitriptravelaxinhotellistgetAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetPageSize is PageSize Setter
// 页大小
func (r *TaobaoalitriptravelaxinhotellistgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoalitriptravelaxinhotellistgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaoalitriptravelaxinhotellistgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoalitriptravelaxinhotellistgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
