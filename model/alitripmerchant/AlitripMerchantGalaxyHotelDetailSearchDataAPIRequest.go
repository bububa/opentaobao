package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyHotelDetailSearchDataAPIRequest 星河-酒店详情页信息获取(新改版) API请求
// alitrip.merchant.galaxy.hotel.detail.search.data
//
// 星河服务，获取雅高酒店详细信息，详情页新改版接口
type AlitripMerchantGalaxyHotelDetailSearchDataAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 请求参数
	_hotelDetailsParam *HotelDetailsParam
}

// NewAlitripMerchantGalaxyHotelDetailSearchDataRequest 初始化AlitripMerchantGalaxyHotelDetailSearchDataAPIRequest对象
func NewAlitripMerchantGalaxyHotelDetailSearchDataRequest() *AlitripMerchantGalaxyHotelDetailSearchDataAPIRequest {
	return &AlitripMerchantGalaxyHotelDetailSearchDataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyHotelDetailSearchDataAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.hotel.detail.search.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyHotelDetailSearchDataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyHotelDetailSearchDataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyHotelDetailSearchDataAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyHotelDetailSearchDataAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetHotelDetailsParam is HotelDetailsParam Setter
// 请求参数
func (r *AlitripMerchantGalaxyHotelDetailSearchDataAPIRequest) SetHotelDetailsParam(_hotelDetailsParam *HotelDetailsParam) error {
	r._hotelDetailsParam = _hotelDetailsParam
	r.Set("hotel_details_param", _hotelDetailsParam)
	return nil
}

// GetHotelDetailsParam HotelDetailsParam Getter
func (r AlitripMerchantGalaxyHotelDetailSearchDataAPIRequest) GetHotelDetailsParam() *HotelDetailsParam {
	return r._hotelDetailsParam
}
