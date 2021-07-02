package icbuproduct

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductInventoryUpdateAPIRequest icbu商品库存更新 API请求
// alibaba.icbu.product.inventory.update
//
// 更新库存信息
type AlibabaIcbuProductInventoryUpdateAPIRequest struct {
	model.Params
	// 更新请求
	_requestParam *ProductInventoryRequest
}

// NewAlibabaIcbuProductInventoryUpdateRequest 初始化AlibabaIcbuProductInventoryUpdateAPIRequest对象
func NewAlibabaIcbuProductInventoryUpdateRequest() *AlibabaIcbuProductInventoryUpdateAPIRequest {
	return &AlibabaIcbuProductInventoryUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductInventoryUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.inventory.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductInventoryUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RequestParam Setter
// 更新请求
func (r *AlibabaIcbuProductInventoryUpdateAPIRequest) SetRequestParam(_requestParam *ProductInventoryRequest) error {
	r._requestParam = _requestParam
	r.Set("request_param", _requestParam)
	return nil
}

// Get RequestParam Getter
func (r AlibabaIcbuProductInventoryUpdateAPIRequest) GetRequestParam() *ProductInventoryRequest {
	return r._requestParam
}
