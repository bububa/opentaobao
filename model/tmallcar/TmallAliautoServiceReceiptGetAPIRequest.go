package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoServiceReceiptGetAPIRequest isv查询服务工单详情 API请求
// tmall.aliauto.service.receipt.get
//
// isv查询服务工单详情
type TmallAliautoServiceReceiptGetAPIRequest struct {
	model.Params
	// 工单号
	_receiptId int64
}

// NewTmallAliautoServiceReceiptGetRequest 初始化TmallAliautoServiceReceiptGetAPIRequest对象
func NewTmallAliautoServiceReceiptGetRequest() *TmallAliautoServiceReceiptGetAPIRequest {
	return &TmallAliautoServiceReceiptGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAliautoServiceReceiptGetAPIRequest) Reset() {
	r._receiptId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoServiceReceiptGetAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.service.receipt.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoServiceReceiptGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoServiceReceiptGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReceiptId is ReceiptId Setter
// 工单号
func (r *TmallAliautoServiceReceiptGetAPIRequest) SetReceiptId(_receiptId int64) error {
	r._receiptId = _receiptId
	r.Set("receipt_id", _receiptId)
	return nil
}

// GetReceiptId ReceiptId Getter
func (r TmallAliautoServiceReceiptGetAPIRequest) GetReceiptId() int64 {
	return r._receiptId
}

var poolTmallAliautoServiceReceiptGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAliautoServiceReceiptGetRequest()
	},
}

// GetTmallAliautoServiceReceiptGetRequest 从 sync.Pool 获取 TmallAliautoServiceReceiptGetAPIRequest
func GetTmallAliautoServiceReceiptGetAPIRequest() *TmallAliautoServiceReceiptGetAPIRequest {
	return poolTmallAliautoServiceReceiptGetAPIRequest.Get().(*TmallAliautoServiceReceiptGetAPIRequest)
}

// ReleaseTmallAliautoServiceReceiptGetAPIRequest 将 TmallAliautoServiceReceiptGetAPIRequest 放入 sync.Pool
func ReleaseTmallAliautoServiceReceiptGetAPIRequest(v *TmallAliautoServiceReceiptGetAPIRequest) {
	v.Reset()
	poolTmallAliautoServiceReceiptGetAPIRequest.Put(v)
}
