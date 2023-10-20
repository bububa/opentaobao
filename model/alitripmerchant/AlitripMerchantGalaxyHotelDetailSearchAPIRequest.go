package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyHotelDetailSearchAPIRequest 星河-酒店详细信息搜索 API请求
// alitrip.merchant.galaxy.hotel.detail.search
//
// 星河服务=获取雅高酒店详细信息
type AlitripMerchantGalaxyHotelDetailSearchAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 酒店详情入参
	_hotelDetailsParam *HotelDetailsParam
}

// NewAlitripMerchantGalaxyHotelDetailSearchRequest 初始化AlitripMerchantGalaxyHotelDetailSearchAPIRequest对象
func NewAlitripMerchantGalaxyHotelDetailSearchRequest() *AlitripMerchantGalaxyHotelDetailSearchAPIRequest {
	return &AlitripMerchantGalaxyHotelDetailSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyHotelDetailSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.hotel.detail.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyHotelDetailSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyHotelDetailSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyHotelDetailSearchAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyHotelDetailSearchAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetHotelDetailsParam is HotelDetailsParam Setter
// 酒店详情入参
func (r *AlitripMerchantGalaxyHotelDetailSearchAPIRequest) SetHotelDetailsParam(_hotelDetailsParam *HotelDetailsParam) error {
	r._hotelDetailsParam = _hotelDetailsParam
	r.Set("hotel_details_param", _hotelDetailsParam)
	return nil
}

// GetHotelDetailsParam HotelDetailsParam Getter
func (r AlitripMerchantGalaxyHotelDetailSearchAPIRequest) GetHotelDetailsParam() *HotelDetailsParam {
	return r._hotelDetailsParam
}
