package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenReturnorderConfirmAPIRequest 退货入库单确认接口 API请求
// taobao.qimen.returnorder.confirm
//
// taobao.qimen.returnorder.confirm
type TaobaoQimenReturnorderConfirmAPIRequest struct {
	model.Params
	//
	_request *ReturnOrderConfirmRequest
}

// NewTaobaoQimenReturnorderConfirmRequest 初始化TaobaoQimenReturnorderConfirmAPIRequest对象
func NewTaobaoQimenReturnorderConfirmRequest() *TaobaoQimenReturnorderConfirmAPIRequest {
	return &TaobaoQimenReturnorderConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenReturnorderConfirmAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenReturnorderConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.returnorder.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenReturnorderConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenReturnorderConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenReturnorderConfirmAPIRequest) SetRequest(_request *ReturnOrderConfirmRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenReturnorderConfirmAPIRequest) GetRequest() *ReturnOrderConfirmRequest {
	return r._request
}

var poolTaobaoQimenReturnorderConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenReturnorderConfirmRequest()
	},
}

// GetTaobaoQimenReturnorderConfirmRequest 从 sync.Pool 获取 TaobaoQimenReturnorderConfirmAPIRequest
func GetTaobaoQimenReturnorderConfirmAPIRequest() *TaobaoQimenReturnorderConfirmAPIRequest {
	return poolTaobaoQimenReturnorderConfirmAPIRequest.Get().(*TaobaoQimenReturnorderConfirmAPIRequest)
}

// ReleaseTaobaoQimenReturnorderConfirmAPIRequest 将 TaobaoQimenReturnorderConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenReturnorderConfirmAPIRequest(v *TaobaoQimenReturnorderConfirmAPIRequest) {
	v.Reset()
	poolTaobaoQimenReturnorderConfirmAPIRequest.Put(v)
}
