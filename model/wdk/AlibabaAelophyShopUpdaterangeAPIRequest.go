package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyShopUpdaterangeAPIRequest 更新渠道店销售范围 API请求
// alibaba.aelophy.shop.updaterange
//
// 更新渠道店销售范围
type AlibabaAelophyShopUpdaterangeAPIRequest struct {
	model.Params
	// 请求对象
	_shopRangeUpdateRequest *ShopRangeUpdateRequest
}

// NewAlibabaAelophyShopUpdaterangeRequest 初始化AlibabaAelophyShopUpdaterangeAPIRequest对象
func NewAlibabaAelophyShopUpdaterangeRequest() *AlibabaAelophyShopUpdaterangeAPIRequest {
	return &AlibabaAelophyShopUpdaterangeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAelophyShopUpdaterangeAPIRequest) Reset() {
	r._shopRangeUpdateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAelophyShopUpdaterangeAPIRequest) GetApiMethodName() string {
	return "alibaba.aelophy.shop.updaterange"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAelophyShopUpdaterangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAelophyShopUpdaterangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopRangeUpdateRequest is ShopRangeUpdateRequest Setter
// 请求对象
func (r *AlibabaAelophyShopUpdaterangeAPIRequest) SetShopRangeUpdateRequest(_shopRangeUpdateRequest *ShopRangeUpdateRequest) error {
	r._shopRangeUpdateRequest = _shopRangeUpdateRequest
	r.Set("shop_range_update_request", _shopRangeUpdateRequest)
	return nil
}

// GetShopRangeUpdateRequest ShopRangeUpdateRequest Getter
func (r AlibabaAelophyShopUpdaterangeAPIRequest) GetShopRangeUpdateRequest() *ShopRangeUpdateRequest {
	return r._shopRangeUpdateRequest
}

var poolAlibabaAelophyShopUpdaterangeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAelophyShopUpdaterangeRequest()
	},
}

// GetAlibabaAelophyShopUpdaterangeRequest 从 sync.Pool 获取 AlibabaAelophyShopUpdaterangeAPIRequest
func GetAlibabaAelophyShopUpdaterangeAPIRequest() *AlibabaAelophyShopUpdaterangeAPIRequest {
	return poolAlibabaAelophyShopUpdaterangeAPIRequest.Get().(*AlibabaAelophyShopUpdaterangeAPIRequest)
}

// ReleaseAlibabaAelophyShopUpdaterangeAPIRequest 将 AlibabaAelophyShopUpdaterangeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAelophyShopUpdaterangeAPIRequest(v *AlibabaAelophyShopUpdaterangeAPIRequest) {
	v.Reset()
	poolAlibabaAelophyShopUpdaterangeAPIRequest.Put(v)
}
