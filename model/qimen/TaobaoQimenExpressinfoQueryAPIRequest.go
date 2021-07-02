package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenExpressinfoQueryAPIRequest 配送公司信息查询接口 API请求
// taobao.qimen.expressinfo.query
//
// 配送公司信息查询
type TaobaoQimenExpressinfoQueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenExpressinfoQueryRequest
}

// NewTaobaoQimenExpressinfoQueryRequest 初始化TaobaoQimenExpressinfoQueryAPIRequest对象
func NewTaobaoQimenExpressinfoQueryRequest() *TaobaoQimenExpressinfoQueryAPIRequest {
	return &TaobaoQimenExpressinfoQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenExpressinfoQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.expressinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenExpressinfoQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRequest is Request Setter
//
func (r *TaobaoQimenExpressinfoQueryAPIRequest) SetRequest(_request *TaobaoQimenExpressinfoQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenExpressinfoQueryAPIRequest) GetRequest() *TaobaoQimenExpressinfoQueryRequest {
	return r._request
}
