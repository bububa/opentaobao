package icbuproduct

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuProductInventoryUpdateAPIRequest) Reset() {
	r._requestParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductInventoryUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.inventory.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductInventoryUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuProductInventoryUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestParam is RequestParam Setter
// 更新请求
func (r *AlibabaIcbuProductInventoryUpdateAPIRequest) SetRequestParam(_requestParam *ProductInventoryRequest) error {
	r._requestParam = _requestParam
	r.Set("request_param", _requestParam)
	return nil
}

// GetRequestParam RequestParam Getter
func (r AlibabaIcbuProductInventoryUpdateAPIRequest) GetRequestParam() *ProductInventoryRequest {
	return r._requestParam
}

var poolAlibabaIcbuProductInventoryUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuProductInventoryUpdateRequest()
	},
}

// GetAlibabaIcbuProductInventoryUpdateRequest 从 sync.Pool 获取 AlibabaIcbuProductInventoryUpdateAPIRequest
func GetAlibabaIcbuProductInventoryUpdateAPIRequest() *AlibabaIcbuProductInventoryUpdateAPIRequest {
	return poolAlibabaIcbuProductInventoryUpdateAPIRequest.Get().(*AlibabaIcbuProductInventoryUpdateAPIRequest)
}

// ReleaseAlibabaIcbuProductInventoryUpdateAPIRequest 将 AlibabaIcbuProductInventoryUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuProductInventoryUpdateAPIRequest(v *AlibabaIcbuProductInventoryUpdateAPIRequest) {
	v.Reset()
	poolAlibabaIcbuProductInventoryUpdateAPIRequest.Put(v)
}
