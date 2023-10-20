package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabagpuupdateschemagetAPIRequest 获取产品编辑schema规则的接口 API请求
// alibaba.gpu.update.schema.get
//
// 获取产品编辑schema规则的接口
type AlibabagpuupdateschemagetAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
	// 当前用户所在渠道如0代表天猫，8代表淘宝
	_providerId int64
}

// NewAlibabagpuupdateschemagetRequest 初始化AlibabagpuupdateschemagetAPIRequest对象
func NewAlibabagpuupdateschemagetRequest() *AlibabagpuupdateschemagetAPIRequest {
	return &AlibabagpuupdateschemagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabagpuupdateschemagetAPIRequest) GetApiMethodName() string {
	return "alibaba.gpu.update.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabagpuupdateschemagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabagpuupdateschemagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品ID
func (r *AlibabagpuupdateschemagetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabagpuupdateschemagetAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetProviderId is ProviderId Setter
// 当前用户所在渠道如0代表天猫，8代表淘宝
func (r *AlibabagpuupdateschemagetAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlibabagpuupdateschemagetAPIRequest) GetProviderId() int64 {
	return r._providerId
}
