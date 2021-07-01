package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDentalItemUnbindAPIRequest
ISV解绑商品 API请求
alibaba.alihealth.dental.item.unbind

ISV解绑商品 */
type AlibabaAlihealthDentalItemUnbindAPIRequest struct {
	model.Params
	// ISV门店ID
	_spStoreId string
	// ISV商品ID
	_spItemId string
}

// NewAlibabaAlihealthDentalItemUnbindRequest 初始化AlibabaAlihealthDentalItemUnbindAPIRequest对象
func NewAlibabaAlihealthDentalItemUnbindRequest() *AlibabaAlihealthDentalItemUnbindAPIRequest {
	return &AlibabaAlihealthDentalItemUnbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalItemUnbindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.item.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalItemUnbindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SpStoreId Setter
// ISV门店ID
func (r *AlibabaAlihealthDentalItemUnbindAPIRequest) SetSpStoreId(_spStoreId string) error {
	r._spStoreId = _spStoreId
	r.Set("sp_store_id", _spStoreId)
	return nil
}

// Get SpStoreId Getter
func (r AlibabaAlihealthDentalItemUnbindAPIRequest) GetSpStoreId() string {
	return r._spStoreId
}

// Set is SpItemId Setter
// ISV商品ID
func (r *AlibabaAlihealthDentalItemUnbindAPIRequest) SetSpItemId(_spItemId string) error {
	r._spItemId = _spItemId
	r.Set("sp_item_id", _spItemId)
	return nil
}

// Get SpItemId Getter
func (r AlibabaAlihealthDentalItemUnbindAPIRequest) GetSpItemId() string {
	return r._spItemId
}
