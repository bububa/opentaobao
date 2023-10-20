package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenItemmappingCreateAPIRequest 前后端商品映射接口 API请求
// taobao.qimen.itemmapping.create
//
// 前后端商品映射
type TaobaoQimenItemmappingCreateAPIRequest struct {
	model.Params
	//
	_request *RequestDo
}

// NewTaobaoQimenItemmappingCreateRequest 初始化TaobaoQimenItemmappingCreateAPIRequest对象
func NewTaobaoQimenItemmappingCreateRequest() *TaobaoQimenItemmappingCreateAPIRequest {
	return &TaobaoQimenItemmappingCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenItemmappingCreateAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemmappingCreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.itemmapping.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemmappingCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenItemmappingCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenItemmappingCreateAPIRequest) SetRequest(_request *RequestDo) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenItemmappingCreateAPIRequest) GetRequest() *RequestDo {
	return r._request
}

var poolTaobaoQimenItemmappingCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenItemmappingCreateRequest()
	},
}

// GetTaobaoQimenItemmappingCreateRequest 从 sync.Pool 获取 TaobaoQimenItemmappingCreateAPIRequest
func GetTaobaoQimenItemmappingCreateAPIRequest() *TaobaoQimenItemmappingCreateAPIRequest {
	return poolTaobaoQimenItemmappingCreateAPIRequest.Get().(*TaobaoQimenItemmappingCreateAPIRequest)
}

// ReleaseTaobaoQimenItemmappingCreateAPIRequest 将 TaobaoQimenItemmappingCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenItemmappingCreateAPIRequest(v *TaobaoQimenItemmappingCreateAPIRequest) {
	v.Reset()
	poolTaobaoQimenItemmappingCreateAPIRequest.Put(v)
}
