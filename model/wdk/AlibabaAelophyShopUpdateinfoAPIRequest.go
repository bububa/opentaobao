package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyShopUpdateinfoAPIRequest 更新渠道店基础信息 API请求
// alibaba.aelophy.shop.updateinfo
//
// 更新渠道店基础信息
type AlibabaAelophyShopUpdateinfoAPIRequest struct {
	model.Params
	// 请求对象
	_shopInfoUpdateRequest *ShopInfoUpdateRequest
}

// NewAlibabaAelophyShopUpdateinfoRequest 初始化AlibabaAelophyShopUpdateinfoAPIRequest对象
func NewAlibabaAelophyShopUpdateinfoRequest() *AlibabaAelophyShopUpdateinfoAPIRequest {
	return &AlibabaAelophyShopUpdateinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAelophyShopUpdateinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.aelophy.shop.updateinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAelophyShopUpdateinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAelophyShopUpdateinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopInfoUpdateRequest is ShopInfoUpdateRequest Setter
// 请求对象
func (r *AlibabaAelophyShopUpdateinfoAPIRequest) SetShopInfoUpdateRequest(_shopInfoUpdateRequest *ShopInfoUpdateRequest) error {
	r._shopInfoUpdateRequest = _shopInfoUpdateRequest
	r.Set("shop_info_update_request", _shopInfoUpdateRequest)
	return nil
}

// GetShopInfoUpdateRequest ShopInfoUpdateRequest Getter
func (r AlibabaAelophyShopUpdateinfoAPIRequest) GetShopInfoUpdateRequest() *ShopInfoUpdateRequest {
	return r._shopInfoUpdateRequest
}
