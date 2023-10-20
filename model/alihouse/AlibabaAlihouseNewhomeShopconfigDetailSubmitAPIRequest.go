package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeshopconfigdetailsubmitAPIRequest 店铺配置信息接口 API请求
// alibaba.alihouse.newhome.shopconfig.detail.submit
//
// 提供店铺配置的能力
type AlibabaalihousenewhomeshopconfigdetailsubmitAPIRequest struct {
	model.Params
	// 城市配置信息
	_shopConfig *ShopConfigDto
}

// NewAlibabaalihousenewhomeshopconfigdetailsubmitRequest 初始化AlibabaalihousenewhomeshopconfigdetailsubmitAPIRequest对象
func NewAlibabaalihousenewhomeshopconfigdetailsubmitRequest() *AlibabaalihousenewhomeshopconfigdetailsubmitAPIRequest {
	return &AlibabaalihousenewhomeshopconfigdetailsubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeshopconfigdetailsubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.shopconfig.detail.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeshopconfigdetailsubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeshopconfigdetailsubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopConfig is ShopConfig Setter
// 城市配置信息
func (r *AlibabaalihousenewhomeshopconfigdetailsubmitAPIRequest) SetShopConfig(_shopConfig *ShopConfigDto) error {
	r._shopConfig = _shopConfig
	r.Set("shop_config", _shopConfig)
	return nil
}

// GetShopConfig ShopConfig Getter
func (r AlibabaalihousenewhomeshopconfigdetailsubmitAPIRequest) GetShopConfig() *ShopConfigDto {
	return r._shopConfig
}
