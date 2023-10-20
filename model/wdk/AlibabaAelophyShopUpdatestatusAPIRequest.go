package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyShopUpdatestatusAPIRequest 更新渠道店营业状态 API请求
// alibaba.aelophy.shop.updatestatus
//
// 更新渠道店营业状态
type AlibabaAelophyShopUpdatestatusAPIRequest struct {
	model.Params
	// 请求对象
	_shopStatusUpdateRequest *ShopStatusUpdateRequest
}

// NewAlibabaAelophyShopUpdatestatusRequest 初始化AlibabaAelophyShopUpdatestatusAPIRequest对象
func NewAlibabaAelophyShopUpdatestatusRequest() *AlibabaAelophyShopUpdatestatusAPIRequest {
	return &AlibabaAelophyShopUpdatestatusAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAelophyShopUpdatestatusAPIRequest) Reset() {
	r._shopStatusUpdateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAelophyShopUpdatestatusAPIRequest) GetApiMethodName() string {
	return "alibaba.aelophy.shop.updatestatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAelophyShopUpdatestatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAelophyShopUpdatestatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopStatusUpdateRequest is ShopStatusUpdateRequest Setter
// 请求对象
func (r *AlibabaAelophyShopUpdatestatusAPIRequest) SetShopStatusUpdateRequest(_shopStatusUpdateRequest *ShopStatusUpdateRequest) error {
	r._shopStatusUpdateRequest = _shopStatusUpdateRequest
	r.Set("shop_status_update_request", _shopStatusUpdateRequest)
	return nil
}

// GetShopStatusUpdateRequest ShopStatusUpdateRequest Getter
func (r AlibabaAelophyShopUpdatestatusAPIRequest) GetShopStatusUpdateRequest() *ShopStatusUpdateRequest {
	return r._shopStatusUpdateRequest
}

var poolAlibabaAelophyShopUpdatestatusAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAelophyShopUpdatestatusRequest()
	},
}

// GetAlibabaAelophyShopUpdatestatusRequest 从 sync.Pool 获取 AlibabaAelophyShopUpdatestatusAPIRequest
func GetAlibabaAelophyShopUpdatestatusAPIRequest() *AlibabaAelophyShopUpdatestatusAPIRequest {
	return poolAlibabaAelophyShopUpdatestatusAPIRequest.Get().(*AlibabaAelophyShopUpdatestatusAPIRequest)
}

// ReleaseAlibabaAelophyShopUpdatestatusAPIRequest 将 AlibabaAelophyShopUpdatestatusAPIRequest 放入 sync.Pool
func ReleaseAlibabaAelophyShopUpdatestatusAPIRequest(v *AlibabaAelophyShopUpdatestatusAPIRequest) {
	v.Reset()
	poolAlibabaAelophyShopUpdatestatusAPIRequest.Put(v)
}
