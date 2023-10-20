package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenChannelinventoryQueryAPIRequest 渠道库存查询接口 API请求
// taobao.qimen.channelinventory.query
//
// 渠道库存查询
type TaobaoQimenChannelinventoryQueryAPIRequest struct {
	model.Params
	//
	_request *RequestDo
}

// NewTaobaoQimenChannelinventoryQueryRequest 初始化TaobaoQimenChannelinventoryQueryAPIRequest对象
func NewTaobaoQimenChannelinventoryQueryRequest() *TaobaoQimenChannelinventoryQueryAPIRequest {
	return &TaobaoQimenChannelinventoryQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenChannelinventoryQueryAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenChannelinventoryQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.channelinventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenChannelinventoryQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenChannelinventoryQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenChannelinventoryQueryAPIRequest) SetRequest(_request *RequestDo) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenChannelinventoryQueryAPIRequest) GetRequest() *RequestDo {
	return r._request
}

var poolTaobaoQimenChannelinventoryQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenChannelinventoryQueryRequest()
	},
}

// GetTaobaoQimenChannelinventoryQueryRequest 从 sync.Pool 获取 TaobaoQimenChannelinventoryQueryAPIRequest
func GetTaobaoQimenChannelinventoryQueryAPIRequest() *TaobaoQimenChannelinventoryQueryAPIRequest {
	return poolTaobaoQimenChannelinventoryQueryAPIRequest.Get().(*TaobaoQimenChannelinventoryQueryAPIRequest)
}

// ReleaseTaobaoQimenChannelinventoryQueryAPIRequest 将 TaobaoQimenChannelinventoryQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenChannelinventoryQueryAPIRequest(v *TaobaoQimenChannelinventoryQueryAPIRequest) {
	v.Reset()
	poolTaobaoQimenChannelinventoryQueryAPIRequest.Put(v)
}
