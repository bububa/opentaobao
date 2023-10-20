package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyhoteldetailsearchAPIRequest 星河-酒店详细信息搜索 API请求
// alitrip.merchant.galaxy.hotel.detail.search
//
// 星河服务=获取雅高酒店详细信息
type AlitripmerchantgalaxyhoteldetailsearchAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 酒店详情入参
	_hotelDetailsParam *HotelDetailsParam
}

// NewAlitripmerchantgalaxyhoteldetailsearchRequest 初始化AlitripmerchantgalaxyhoteldetailsearchAPIRequest对象
func NewAlitripmerchantgalaxyhoteldetailsearchRequest() *AlitripmerchantgalaxyhoteldetailsearchAPIRequest {
	return &AlitripmerchantgalaxyhoteldetailsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyhoteldetailsearchAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.hotel.detail.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyhoteldetailsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyhoteldetailsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripmerchantgalaxyhoteldetailsearchAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyhoteldetailsearchAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetHotelDetailsParam is HotelDetailsParam Setter
// 酒店详情入参
func (r *AlitripmerchantgalaxyhoteldetailsearchAPIRequest) SetHotelDetailsParam(_hotelDetailsParam *HotelDetailsParam) error {
	r._hotelDetailsParam = _hotelDetailsParam
	r.Set("hotel_details_param", _hotelDetailsParam)
	return nil
}

// GetHotelDetailsParam HotelDetailsParam Getter
func (r AlitripmerchantgalaxyhoteldetailsearchAPIRequest) GetHotelDetailsParam() *HotelDetailsParam {
	return r._hotelDetailsParam
}
