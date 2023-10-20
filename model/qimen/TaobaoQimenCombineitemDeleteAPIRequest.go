package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenCombineitemDeleteAPIRequest 组合货品删除接口 API请求
// taobao.qimen.combineitem.delete
//
// 组合货品删除
type TaobaoQimenCombineitemDeleteAPIRequest struct {
	model.Params
	//
	_request *RequestDo
}

// NewTaobaoQimenCombineitemDeleteRequest 初始化TaobaoQimenCombineitemDeleteAPIRequest对象
func NewTaobaoQimenCombineitemDeleteRequest() *TaobaoQimenCombineitemDeleteAPIRequest {
	return &TaobaoQimenCombineitemDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenCombineitemDeleteAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenCombineitemDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.combineitem.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenCombineitemDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenCombineitemDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenCombineitemDeleteAPIRequest) SetRequest(_request *RequestDo) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenCombineitemDeleteAPIRequest) GetRequest() *RequestDo {
	return r._request
}

var poolTaobaoQimenCombineitemDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenCombineitemDeleteRequest()
	},
}

// GetTaobaoQimenCombineitemDeleteRequest 从 sync.Pool 获取 TaobaoQimenCombineitemDeleteAPIRequest
func GetTaobaoQimenCombineitemDeleteAPIRequest() *TaobaoQimenCombineitemDeleteAPIRequest {
	return poolTaobaoQimenCombineitemDeleteAPIRequest.Get().(*TaobaoQimenCombineitemDeleteAPIRequest)
}

// ReleaseTaobaoQimenCombineitemDeleteAPIRequest 将 TaobaoQimenCombineitemDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenCombineitemDeleteAPIRequest(v *TaobaoQimenCombineitemDeleteAPIRequest) {
	v.Reset()
	poolTaobaoQimenCombineitemDeleteAPIRequest.Put(v)
}
