package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaelophyshopupdatestatusAPIRequest 更新渠道店营业状态 API请求
// alibaba.aelophy.shop.updatestatus
//
// 更新渠道店营业状态
type AlibabaaelophyshopupdatestatusAPIRequest struct {
	model.Params
	// 请求对象
	_shopStatusUpdateRequest *ShopStatusUpdateRequest
}

// NewAlibabaaelophyshopupdatestatusRequest 初始化AlibabaaelophyshopupdatestatusAPIRequest对象
func NewAlibabaaelophyshopupdatestatusRequest() *AlibabaaelophyshopupdatestatusAPIRequest {
	return &AlibabaaelophyshopupdatestatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaelophyshopupdatestatusAPIRequest) GetApiMethodName() string {
	return "alibaba.aelophy.shop.updatestatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaelophyshopupdatestatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaelophyshopupdatestatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopStatusUpdateRequest is ShopStatusUpdateRequest Setter
// 请求对象
func (r *AlibabaaelophyshopupdatestatusAPIRequest) SetShopStatusUpdateRequest(_shopStatusUpdateRequest *ShopStatusUpdateRequest) error {
	r._shopStatusUpdateRequest = _shopStatusUpdateRequest
	r.Set("shop_status_update_request", _shopStatusUpdateRequest)
	return nil
}

// GetShopStatusUpdateRequest ShopStatusUpdateRequest Getter
func (r AlibabaaelophyshopupdatestatusAPIRequest) GetShopStatusUpdateRequest() *ShopStatusUpdateRequest {
	return r._shopStatusUpdateRequest
}
