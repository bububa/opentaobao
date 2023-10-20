package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderQueryAPIRequest 根据收件人信息查询交易单号接口 API请求
// taobao.qimen.order.query
//
// WMS 调用该接口，根据收件人信息查询平台交易订单号。
type TaobaoQimenOrderQueryAPIRequest struct {
	model.Params
	// request
	_request *TaobaoQimenOrderQueryRequest
}

// NewTaobaoQimenOrderQueryRequest 初始化TaobaoQimenOrderQueryAPIRequest对象
func NewTaobaoQimenOrderQueryRequest() *TaobaoQimenOrderQueryAPIRequest {
	return &TaobaoQimenOrderQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenOrderQueryAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
// request
func (r *TaobaoQimenOrderQueryAPIRequest) SetRequest(_request *TaobaoQimenOrderQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenOrderQueryAPIRequest) GetRequest() *TaobaoQimenOrderQueryRequest {
	return r._request
}

var poolTaobaoQimenOrderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenOrderQueryRequest()
	},
}

// GetTaobaoQimenOrderQueryRequest 从 sync.Pool 获取 TaobaoQimenOrderQueryAPIRequest
func GetTaobaoQimenOrderQueryAPIRequest() *TaobaoQimenOrderQueryAPIRequest {
	return poolTaobaoQimenOrderQueryAPIRequest.Get().(*TaobaoQimenOrderQueryAPIRequest)
}

// ReleaseTaobaoQimenOrderQueryAPIRequest 将 TaobaoQimenOrderQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenOrderQueryAPIRequest(v *TaobaoQimenOrderQueryAPIRequest) {
	v.Reset()
	poolTaobaoQimenOrderQueryAPIRequest.Put(v)
}
