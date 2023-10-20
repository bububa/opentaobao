package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimeninventoryrulecreateAPIRequest 渠道间库存规则设置接口 API请求
// taobao.qimen.inventoryrule.create
//
// 渠道间库存规则设置
type TaobaoqimeninventoryrulecreateAPIRequest struct {
	model.Params
	//
	_request *RequestDo
}

// NewTaobaoqimeninventoryrulecreateRequest 初始化TaobaoqimeninventoryrulecreateAPIRequest对象
func NewTaobaoqimeninventoryrulecreateRequest() *TaobaoqimeninventoryrulecreateAPIRequest {
	return &TaobaoqimeninventoryrulecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimeninventoryrulecreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventoryrule.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimeninventoryrulecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimeninventoryrulecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimeninventoryrulecreateAPIRequest) SetRequest(_request *RequestDo) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimeninventoryrulecreateAPIRequest) GetRequest() *RequestDo {
	return r._request
}
