package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest 店铺配置信息接口 API请求
// alibaba.alihouse.newhome.shopconfig.detail.submit
//
// 提供店铺配置的能力
type AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest struct {
	model.Params
	// 城市配置信息
	_shopConfig *ShopConfigDto
}

// NewAlibabaAlihouseNewhomeShopconfigDetailSubmitRequest 初始化AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeShopconfigDetailSubmitRequest() *AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest) Reset() {
	r._shopConfig = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.shopconfig.detail.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopConfig is ShopConfig Setter
// 城市配置信息
func (r *AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest) SetShopConfig(_shopConfig *ShopConfigDto) error {
	r._shopConfig = _shopConfig
	r.Set("shop_config", _shopConfig)
	return nil
}

// GetShopConfig ShopConfig Getter
func (r AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest) GetShopConfig() *ShopConfigDto {
	return r._shopConfig
}

var poolAlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeShopconfigDetailSubmitRequest()
	},
}

// GetAlibabaAlihouseNewhomeShopconfigDetailSubmitRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest
func GetAlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest() *AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest {
	return poolAlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest.Get().(*AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest 将 AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest(v *AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest.Put(v)
}
