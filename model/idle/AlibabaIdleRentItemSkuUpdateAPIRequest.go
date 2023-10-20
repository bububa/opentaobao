package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerentitemskuupdateAPIRequest 更新/增加sku信息 API请求
// alibaba.idle.rent.item.sku.update
//
// 更新/增加sku信息
type AlibabaidlerentitemskuupdateAPIRequest struct {
	model.Params
	// sku信息，更新后skuId保持不变
	_sku *ItemSkuDto
}

// NewAlibabaidlerentitemskuupdateRequest 初始化AlibabaidlerentitemskuupdateAPIRequest对象
func NewAlibabaidlerentitemskuupdateRequest() *AlibabaidlerentitemskuupdateAPIRequest {
	return &AlibabaidlerentitemskuupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlerentitemskuupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.item.sku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlerentitemskuupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlerentitemskuupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSku is Sku Setter
// sku信息，更新后skuId保持不变
func (r *AlibabaidlerentitemskuupdateAPIRequest) SetSku(_sku *ItemSkuDto) error {
	r._sku = _sku
	r.Set("sku", _sku)
	return nil
}

// GetSku Sku Getter
func (r AlibabaidlerentitemskuupdateAPIRequest) GetSku() *ItemSkuDto {
	return r._sku
}
