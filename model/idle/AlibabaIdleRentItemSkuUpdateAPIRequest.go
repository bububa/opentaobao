package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentItemSkuUpdateAPIRequest 更新/增加sku信息 API请求
// alibaba.idle.rent.item.sku.update
//
// 更新/增加sku信息
type AlibabaIdleRentItemSkuUpdateAPIRequest struct {
	model.Params
	// sku信息，更新后skuId保持不变
	_sku *ItemSkuDto
}

// NewAlibabaIdleRentItemSkuUpdateRequest 初始化AlibabaIdleRentItemSkuUpdateAPIRequest对象
func NewAlibabaIdleRentItemSkuUpdateRequest() *AlibabaIdleRentItemSkuUpdateAPIRequest {
	return &AlibabaIdleRentItemSkuUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleRentItemSkuUpdateAPIRequest) Reset() {
	r._sku = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentItemSkuUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.item.sku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentItemSkuUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleRentItemSkuUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSku is Sku Setter
// sku信息，更新后skuId保持不变
func (r *AlibabaIdleRentItemSkuUpdateAPIRequest) SetSku(_sku *ItemSkuDto) error {
	r._sku = _sku
	r.Set("sku", _sku)
	return nil
}

// GetSku Sku Getter
func (r AlibabaIdleRentItemSkuUpdateAPIRequest) GetSku() *ItemSkuDto {
	return r._sku
}

var poolAlibabaIdleRentItemSkuUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleRentItemSkuUpdateRequest()
	},
}

// GetAlibabaIdleRentItemSkuUpdateRequest 从 sync.Pool 获取 AlibabaIdleRentItemSkuUpdateAPIRequest
func GetAlibabaIdleRentItemSkuUpdateAPIRequest() *AlibabaIdleRentItemSkuUpdateAPIRequest {
	return poolAlibabaIdleRentItemSkuUpdateAPIRequest.Get().(*AlibabaIdleRentItemSkuUpdateAPIRequest)
}

// ReleaseAlibabaIdleRentItemSkuUpdateAPIRequest 将 AlibabaIdleRentItemSkuUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleRentItemSkuUpdateAPIRequest(v *AlibabaIdleRentItemSkuUpdateAPIRequest) {
	v.Reset()
	poolAlibabaIdleRentItemSkuUpdateAPIRequest.Put(v)
}
