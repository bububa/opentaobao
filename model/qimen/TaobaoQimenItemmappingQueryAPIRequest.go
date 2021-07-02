package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenItemmappingQueryAPIRequest 前后端商品映射查询接口 API请求
// taobao.qimen.itemmapping.query
//
// 前后端商品映射查询接口
type TaobaoQimenItemmappingQueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenItemmappingQueryRequest
}

// NewTaobaoQimenItemmappingQueryRequest 初始化TaobaoQimenItemmappingQueryAPIRequest对象
func NewTaobaoQimenItemmappingQueryRequest() *TaobaoQimenItemmappingQueryAPIRequest {
	return &TaobaoQimenItemmappingQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemmappingQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.itemmapping.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemmappingQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRequest is Request Setter
//
func (r *TaobaoQimenItemmappingQueryAPIRequest) SetRequest(_request *TaobaoQimenItemmappingQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenItemmappingQueryAPIRequest) GetRequest() *TaobaoQimenItemmappingQueryRequest {
	return r._request
}
