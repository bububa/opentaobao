package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIRequest 城市配置信息接口 API请求
// alibaba.alihouse.newhome.shopcityconfig.detail.submit
//
// 上传城市配置信息
type AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIRequest struct {
	model.Params
	// 城市配置信息
	_shopCityConfig *ShopCityConfigDto
}

// NewAlibabaAlihouseNewhomeShopcityconfigDetailSubmitRequest 初始化AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeShopcityconfigDetailSubmitRequest() *AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.shopcityconfig.detail.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopCityConfig is ShopCityConfig Setter
// 城市配置信息
func (r *AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIRequest) SetShopCityConfig(_shopCityConfig *ShopCityConfigDto) error {
	r._shopCityConfig = _shopCityConfig
	r.Set("shop_city_config", _shopCityConfig)
	return nil
}

// GetShopCityConfig ShopCityConfig Getter
func (r AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIRequest) GetShopCityConfig() *ShopCityConfigDto {
	return r._shopCityConfig
}
