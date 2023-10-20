package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaonetouchlogisticsexpressspecialproducttypelistAPIRequest 获取商品类型配置项 API请求
// alibaba.onetouch.logistics.express.special.product.type.list
//
// 获取商品类型配置项
type AlibabaonetouchlogisticsexpressspecialproducttypelistAPIRequest struct {
	model.Params
}

// NewAlibabaonetouchlogisticsexpressspecialproducttypelistRequest 初始化AlibabaonetouchlogisticsexpressspecialproducttypelistAPIRequest对象
func NewAlibabaonetouchlogisticsexpressspecialproducttypelistRequest() *AlibabaonetouchlogisticsexpressspecialproducttypelistAPIRequest {
	return &AlibabaonetouchlogisticsexpressspecialproducttypelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaonetouchlogisticsexpressspecialproducttypelistAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.special.product.type.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaonetouchlogisticsexpressspecialproducttypelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaonetouchlogisticsexpressspecialproducttypelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}
