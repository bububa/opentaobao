package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest 疫苗预约门店列表查询 API请求
// alibaba.taobao.micdetail.alihealth.querystores
//
// 微信小程序疫苗预约门店列表查询
type AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest struct {
	model.Params
	// 0不查询库存1查询库存
	_invType int64
	// 包ID
	_packageId int64
	// 页大小
	_pageSize int64
	// 页码
	_pageNum int64
	// 经度
	_lon string
	// 维度
	_lat string
	// 商品ID
	_itemId string
	// 用户ID
	_userId int64
	// 地区ID
	_districtId int64
	// 城市ID
	_cityId int64
	// 省份ID
	_provinceId int64
}

// NewAlibabaTaobaoMicdetailAlihealthQuerystoresRequest 初始化AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest对象
func NewAlibabaTaobaoMicdetailAlihealthQuerystoresRequest() *AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest {
	return &AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) GetApiMethodName() string {
	return "alibaba.taobao.micdetail.alihealth.querystores"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetInvType is InvType Setter
// 0不查询库存1查询库存
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) SetInvType(_invType int64) error {
	r._invType = _invType
	r.Set("inv_type", _invType)
	return nil
}

// GetInvType InvType Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) GetInvType() int64 {
	return r._invType
}

// SetPackageId is PackageId Setter
// 包ID
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) SetPackageId(_packageId int64) error {
	r._packageId = _packageId
	r.Set("package_id", _packageId)
	return nil
}

// GetPackageId PackageId Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) GetPackageId() int64 {
	return r._packageId
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 页码
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetLon is Lon Setter
// 经度
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) SetLon(_lon string) error {
	r._lon = _lon
	r.Set("lon", _lon)
	return nil
}

// GetLon Lon Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) GetLon() string {
	return r._lon
}

// SetLat is Lat Setter
// 维度
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) SetLat(_lat string) error {
	r._lat = _lat
	r.Set("lat", _lat)
	return nil
}

// GetLat Lat Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) GetLat() string {
	return r._lat
}

// SetItemId is ItemId Setter
// 商品ID
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) GetItemId() string {
	return r._itemId
}

// SetUserId is UserId Setter
// 用户ID
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) GetUserId() int64 {
	return r._userId
}

// SetDistrictId is DistrictId Setter
// 地区ID
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) SetDistrictId(_districtId int64) error {
	r._districtId = _districtId
	r.Set("district_id", _districtId)
	return nil
}

// GetDistrictId DistrictId Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) GetDistrictId() int64 {
	return r._districtId
}

// SetCityId is CityId Setter
// 城市ID
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) SetCityId(_cityId int64) error {
	r._cityId = _cityId
	r.Set("city_id", _cityId)
	return nil
}

// GetCityId CityId Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) GetCityId() int64 {
	return r._cityId
}

// SetProvinceId is ProvinceId Setter
// 省份ID
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) SetProvinceId(_provinceId int64) error {
	r._provinceId = _provinceId
	r.Set("province_id", _provinceId)
	return nil
}

// GetProvinceId ProvinceId Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest) GetProvinceId() int64 {
	return r._provinceId
}
