package qimen

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenCombineitemQueryAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenCombineitemQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.combineitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenCombineitemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenCombineitemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenCombineitemQueryAPIRequest) SetRequest(_request *TaobaoQimenCombineitemQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenCombineitemQueryAPIRequest) GetRequest() *TaobaoQimenCombineitemQueryRequest {
	return r._request
}

var poolTaobaoQimenCombineitemQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenCombineitemQueryRequest()
	},
}

// GetTaobaoQimenCombineitemQueryRequest 从 sync.Pool 获取 TaobaoQimenCombineitemQueryAPIRequest
func GetTaobaoQimenCombineitemQueryAPIRequest() *TaobaoQimenCombineitemQueryAPIRequest {
	return poolTaobaoQimenCombineitemQueryAPIRequest.Get().(*TaobaoQimenCombineitemQueryAPIRequest)
}

// ReleaseTaobaoQimenCombineitemQueryAPIRequest 将 TaobaoQimenCombineitemQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenCombineitemQueryAPIRequest(v *TaobaoQimenCombineitemQueryAPIRequest) {
	v.Reset()
	poolTaobaoQimenCombineitemQueryAPIRequest.Put(v)
}
