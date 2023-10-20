package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskubarcodequeryAPIRequest 商品条码查询接口 API请求
// alibaba.wdk.sku.barcode.query
//
// 查询商品编码，支持一品多码
type AlibabawdkskubarcodequeryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
}

// NewAlibabawdkskubarcodequeryRequest 初始化AlibabawdkskubarcodequeryAPIRequest对象
func NewAlibabawdkskubarcodequeryRequest() *AlibabawdkskubarcodequeryAPIRequest {
	return &AlibabawdkskubarcodequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskubarcodequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.barcode.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskubarcodequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskubarcodequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabawdkskubarcodequeryAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabawdkskubarcodequeryAPIRequest) GetSkuCode() string {
	return r._skuCode
}
