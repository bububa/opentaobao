package qimen

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenExpressinfoQueryAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenExpressinfoQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.expressinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenExpressinfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenExpressinfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenExpressinfoQueryAPIRequest) SetRequest(_request *TaobaoQimenExpressinfoQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenExpressinfoQueryAPIRequest) GetRequest() *TaobaoQimenExpressinfoQueryRequest {
	return r._request
}

var poolTaobaoQimenExpressinfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenExpressinfoQueryRequest()
	},
}

// GetTaobaoQimenExpressinfoQueryRequest 从 sync.Pool 获取 TaobaoQimenExpressinfoQueryAPIRequest
func GetTaobaoQimenExpressinfoQueryAPIRequest() *TaobaoQimenExpressinfoQueryAPIRequest {
	return poolTaobaoQimenExpressinfoQueryAPIRequest.Get().(*TaobaoQimenExpressinfoQueryAPIRequest)
}

// ReleaseTaobaoQimenExpressinfoQueryAPIRequest 将 TaobaoQimenExpressinfoQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenExpressinfoQueryAPIRequest(v *TaobaoQimenExpressinfoQueryAPIRequest) {
	v.Reset()
	poolTaobaoQimenExpressinfoQueryAPIRequest.Put(v)
}
