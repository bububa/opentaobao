package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenShopSynchronizeAPIRequest 店铺同步接口 API请求
// taobao.qimen.shop.synchronize
//
// 店铺同步接口描述
type TaobaoQimenShopSynchronizeAPIRequest struct {
	model.Params
	// 请求
	_request *TaobaoQimenShopSynchronizeRequest
}

// NewTaobaoQimenShopSynchronizeRequest 初始化TaobaoQimenShopSynchronizeAPIRequest对象
func NewTaobaoQimenShopSynchronizeRequest() *TaobaoQimenShopSynchronizeAPIRequest {
	return &TaobaoQimenShopSynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenShopSynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.shop.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenShopSynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenShopSynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
// 请求
func (r *TaobaoQimenShopSynchronizeAPIRequest) SetRequest(_request *TaobaoQimenShopSynchronizeRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenShopSynchronizeAPIRequest) GetRequest() *TaobaoQimenShopSynchronizeRequest {
	return r._request
}
