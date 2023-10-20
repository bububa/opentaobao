package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenTransferorderQueryAPIRequest 调拨单查询 API请求
// taobao.qimen.transferorder.query
//
// 调拨单查询
type TaobaoQimenTransferorderQueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenTransferorderQueryStruct
}

// NewTaobaoQimenTransferorderQueryRequest 初始化TaobaoQimenTransferorderQueryAPIRequest对象
func NewTaobaoQimenTransferorderQueryRequest() *TaobaoQimenTransferorderQueryAPIRequest {
	return &TaobaoQimenTransferorderQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenTransferorderQueryAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenTransferorderQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.transferorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenTransferorderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenTransferorderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenTransferorderQueryAPIRequest) SetRequest(_request *TaobaoQimenTransferorderQueryStruct) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenTransferorderQueryAPIRequest) GetRequest() *TaobaoQimenTransferorderQueryStruct {
	return r._request
}

var poolTaobaoQimenTransferorderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenTransferorderQueryRequest()
	},
}

// GetTaobaoQimenTransferorderQueryRequest 从 sync.Pool 获取 TaobaoQimenTransferorderQueryAPIRequest
func GetTaobaoQimenTransferorderQueryAPIRequest() *TaobaoQimenTransferorderQueryAPIRequest {
	return poolTaobaoQimenTransferorderQueryAPIRequest.Get().(*TaobaoQimenTransferorderQueryAPIRequest)
}

// ReleaseTaobaoQimenTransferorderQueryAPIRequest 将 TaobaoQimenTransferorderQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenTransferorderQueryAPIRequest(v *TaobaoQimenTransferorderQueryAPIRequest) {
	v.Reset()
	poolTaobaoQimenTransferorderQueryAPIRequest.Put(v)
}
