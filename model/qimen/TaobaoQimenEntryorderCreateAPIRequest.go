package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenEntryorderCreateAPIRequest 入库单创建接口 API请求
// taobao.qimen.entryorder.create
//
// taobao.qimen.entryorder.create
type TaobaoQimenEntryorderCreateAPIRequest struct {
	model.Params
	//
	_request *EntryOrderCreateRequest
}

// NewTaobaoQimenEntryorderCreateRequest 初始化TaobaoQimenEntryorderCreateAPIRequest对象
func NewTaobaoQimenEntryorderCreateRequest() *TaobaoQimenEntryorderCreateAPIRequest {
	return &TaobaoQimenEntryorderCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenEntryorderCreateAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenEntryorderCreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.entryorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenEntryorderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenEntryorderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenEntryorderCreateAPIRequest) SetRequest(_request *EntryOrderCreateRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenEntryorderCreateAPIRequest) GetRequest() *EntryOrderCreateRequest {
	return r._request
}

var poolTaobaoQimenEntryorderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenEntryorderCreateRequest()
	},
}

// GetTaobaoQimenEntryorderCreateRequest 从 sync.Pool 获取 TaobaoQimenEntryorderCreateAPIRequest
func GetTaobaoQimenEntryorderCreateAPIRequest() *TaobaoQimenEntryorderCreateAPIRequest {
	return poolTaobaoQimenEntryorderCreateAPIRequest.Get().(*TaobaoQimenEntryorderCreateAPIRequest)
}

// ReleaseTaobaoQimenEntryorderCreateAPIRequest 将 TaobaoQimenEntryorderCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenEntryorderCreateAPIRequest(v *TaobaoQimenEntryorderCreateAPIRequest) {
	v.Reset()
	poolTaobaoQimenEntryorderCreateAPIRequest.Put(v)
}
