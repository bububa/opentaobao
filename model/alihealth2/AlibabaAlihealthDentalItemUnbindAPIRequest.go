package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalItemUnbindAPIRequest ISV解绑商品 API请求
// alibaba.alihealth.dental.item.unbind
//
// ISV解绑商品
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDentalItemUnbindAPIRequest) Reset() {
	r._spStoreId = ""
	r._spItemId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalItemUnbindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.item.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalItemUnbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDentalItemUnbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSpStoreId is SpStoreId Setter
// ISV门店ID
func (r *AlibabaAlihealthDentalItemUnbindAPIRequest) SetSpStoreId(_spStoreId string) error {
	r._spStoreId = _spStoreId
	r.Set("sp_store_id", _spStoreId)
	return nil
}

// GetSpStoreId SpStoreId Getter
func (r AlibabaAlihealthDentalItemUnbindAPIRequest) GetSpStoreId() string {
	return r._spStoreId
}

// SetSpItemId is SpItemId Setter
// ISV商品ID
func (r *AlibabaAlihealthDentalItemUnbindAPIRequest) SetSpItemId(_spItemId string) error {
	r._spItemId = _spItemId
	r.Set("sp_item_id", _spItemId)
	return nil
}

// GetSpItemId SpItemId Getter
func (r AlibabaAlihealthDentalItemUnbindAPIRequest) GetSpItemId() string {
	return r._spItemId
}

var poolAlibabaAlihealthDentalItemUnbindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDentalItemUnbindRequest()
	},
}

// GetAlibabaAlihealthDentalItemUnbindRequest 从 sync.Pool 获取 AlibabaAlihealthDentalItemUnbindAPIRequest
func GetAlibabaAlihealthDentalItemUnbindAPIRequest() *AlibabaAlihealthDentalItemUnbindAPIRequest {
	return poolAlibabaAlihealthDentalItemUnbindAPIRequest.Get().(*AlibabaAlihealthDentalItemUnbindAPIRequest)
}

// ReleaseAlibabaAlihealthDentalItemUnbindAPIRequest 将 AlibabaAlihealthDentalItemUnbindAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDentalItemUnbindAPIRequest(v *AlibabaAlihealthDentalItemUnbindAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDentalItemUnbindAPIRequest.Put(v)
}
