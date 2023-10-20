package icbuproduct

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductinventoryupdateAPIRequest icbu商品库存更新 API请求
// alibaba.icbu.product.inventory.update
//
// 更新库存信息
type AlibabaicbuproductinventoryupdateAPIRequest struct {
	model.Params
	// 更新请求
	_requestParam *ProductInventoryRequest
}

// NewAlibabaicbuproductinventoryupdateRequest 初始化AlibabaicbuproductinventoryupdateAPIRequest对象
func NewAlibabaicbuproductinventoryupdateRequest() *AlibabaicbuproductinventoryupdateAPIRequest {
	return &AlibabaicbuproductinventoryupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuproductinventoryupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.inventory.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuproductinventoryupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuproductinventoryupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestParam is RequestParam Setter
// 更新请求
func (r *AlibabaicbuproductinventoryupdateAPIRequest) SetRequestParam(_requestParam *ProductInventoryRequest) error {
	r._requestParam = _requestParam
	r.Set("request_param", _requestParam)
	return nil
}

// GetRequestParam RequestParam Getter
func (r AlibabaicbuproductinventoryupdateAPIRequest) GetRequestParam() *ProductInventoryRequest {
	return r._requestParam
}
