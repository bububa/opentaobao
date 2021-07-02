package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventoryruleCreateAPIRequest 渠道间库存规则设置接口 API请求
// taobao.qimen.inventoryrule.create
//
// 渠道间库存规则设置
type TaobaoQimenInventoryruleCreateAPIRequest struct {
	model.Params
	//
	_request *RequestDo
}

// NewTaobaoQimenInventoryruleCreateRequest 初始化TaobaoQimenInventoryruleCreateAPIRequest对象
func NewTaobaoQimenInventoryruleCreateRequest() *TaobaoQimenInventoryruleCreateAPIRequest {
	return &TaobaoQimenInventoryruleCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenInventoryruleCreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventoryrule.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenInventoryruleCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRequest is Request Setter
//
func (r *TaobaoQimenInventoryruleCreateAPIRequest) SetRequest(_request *RequestDo) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenInventoryruleCreateAPIRequest) GetRequest() *RequestDo {
	return r._request
}
