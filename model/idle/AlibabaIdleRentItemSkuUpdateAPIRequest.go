package idle

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentItemSkuUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.item.sku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentItemSkuUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Sku Setter
// sku信息，更新后skuId保持不变
func (r *AlibabaIdleRentItemSkuUpdateAPIRequest) SetSku(_sku *ItemSkuDto) error {
	r._sku = _sku
	r.Set("sku", _sku)
	return nil
}

// Get Sku Getter
func (r AlibabaIdleRentItemSkuUpdateAPIRequest) GetSku() *ItemSkuDto {
	return r._sku
}
