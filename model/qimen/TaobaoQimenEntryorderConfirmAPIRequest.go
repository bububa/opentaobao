package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenEntryorderConfirmAPIRequest 入库单确认接口 API请求
// taobao.qimen.entryorder.confirm
//
// WMS调用接口，回传入库单信息;
type TaobaoQimenEntryorderConfirmAPIRequest struct {
	model.Params
	//
	_request *EntryOrderConfirmRequest
}

// NewTaobaoQimenEntryorderConfirmRequest 初始化TaobaoQimenEntryorderConfirmAPIRequest对象
func NewTaobaoQimenEntryorderConfirmRequest() *TaobaoQimenEntryorderConfirmAPIRequest {
	return &TaobaoQimenEntryorderConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenEntryorderConfirmAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenEntryorderConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.entryorder.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenEntryorderConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenEntryorderConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenEntryorderConfirmAPIRequest) SetRequest(_request *EntryOrderConfirmRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenEntryorderConfirmAPIRequest) GetRequest() *EntryOrderConfirmRequest {
	return r._request
}

var poolTaobaoQimenEntryorderConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenEntryorderConfirmRequest()
	},
}

// GetTaobaoQimenEntryorderConfirmRequest 从 sync.Pool 获取 TaobaoQimenEntryorderConfirmAPIRequest
func GetTaobaoQimenEntryorderConfirmAPIRequest() *TaobaoQimenEntryorderConfirmAPIRequest {
	return poolTaobaoQimenEntryorderConfirmAPIRequest.Get().(*TaobaoQimenEntryorderConfirmAPIRequest)
}

// ReleaseTaobaoQimenEntryorderConfirmAPIRequest 将 TaobaoQimenEntryorderConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenEntryorderConfirmAPIRequest(v *TaobaoQimenEntryorderConfirmAPIRequest) {
	v.Reset()
	poolTaobaoQimenEntryorderConfirmAPIRequest.Put(v)
}
