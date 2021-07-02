package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenCombineitemQueryAPIRequest 组合货品关系查询接口 API请求
// taobao.qimen.combineitem.query
//
// 组合货品关系查询
type TaobaoQimenCombineitemQueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenCombineitemQueryRequest
}

// NewTaobaoQimenCombineitemQueryRequest 初始化TaobaoQimenCombineitemQueryAPIRequest对象
func NewTaobaoQimenCombineitemQueryRequest() *TaobaoQimenCombineitemQueryAPIRequest {
	return &TaobaoQimenCombineitemQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenCombineitemQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.combineitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenCombineitemQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
//
func (r *TaobaoQimenCombineitemQueryAPIRequest) SetRequest(_request *TaobaoQimenCombineitemQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenCombineitemQueryAPIRequest) GetRequest() *TaobaoQimenCombineitemQueryRequest {
	return r._request
}
