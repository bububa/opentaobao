package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyHotelListSearchAPIRequest 星河-酒店列表页搜索 API请求
// alitrip.merchant.galaxy.hotel.list.search
//
// 星河产品=酒店列表页搜索
type AlitripMerchantGalaxyHotelListSearchAPIRequest struct {
	model.Params
	// 商家租户id
	_tenantKey string
	// 请求参数
	_listSearchParam *ListSearchParam
}

// NewAlitripMerchantGalaxyHotelListSearchRequest 初始化AlitripMerchantGalaxyHotelListSearchAPIRequest对象
func NewAlitripMerchantGalaxyHotelListSearchRequest() *AlitripMerchantGalaxyHotelListSearchAPIRequest {
	return &AlitripMerchantGalaxyHotelListSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyHotelListSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.hotel.list.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyHotelListSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyHotelListSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 商家租户id
func (r *AlitripMerchantGalaxyHotelListSearchAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyHotelListSearchAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetListSearchParam is ListSearchParam Setter
// 请求参数
func (r *AlitripMerchantGalaxyHotelListSearchAPIRequest) SetListSearchParam(_listSearchParam *ListSearchParam) error {
	r._listSearchParam = _listSearchParam
	r.Set("list_search_param", _listSearchParam)
	return nil
}

// GetListSearchParam ListSearchParam Getter
func (r AlitripMerchantGalaxyHotelListSearchAPIRequest) GetListSearchParam() *ListSearchParam {
	return r._listSearchParam
}
