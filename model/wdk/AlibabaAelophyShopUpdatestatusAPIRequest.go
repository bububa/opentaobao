package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
