package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimencombineitemdeleteAPIRequest 组合货品删除接口 API请求
// taobao.qimen.combineitem.delete
//
// 组合货品删除
type TaobaoqimencombineitemdeleteAPIRequest struct {
	model.Params
	//
	_request *RequestDo
}

// NewTaobaoqimencombineitemdeleteRequest 初始化TaobaoqimencombineitemdeleteAPIRequest对象
func NewTaobaoqimencombineitemdeleteRequest() *TaobaoqimencombineitemdeleteAPIRequest {
	return &TaobaoqimencombineitemdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimencombineitemdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.combineitem.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimencombineitemdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimencombineitemdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimencombineitemdeleteAPIRequest) SetRequest(_request *RequestDo) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimencombineitemdeleteAPIRequest) GetRequest() *RequestDo {
	return r._request
}
