package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkOrderRefundGetAPIRequest 淘宝客-推广者-全量维权退款订单查询 API请求
// taobao.tbk.order.refund.get
//
// 淘宝客账户下全量维权退款订单查询
type TaobaoTbkOrderRefundGetAPIRequest struct {
	model.Params
	// 全量维权订单查询入参
	_publisherRefundOrderQueryOption *PublisherRefundOrderQueryOption
}

// NewTaobaoTbkOrderRefundGetRequest 初始化TaobaoTbkOrderRefundGetAPIRequest对象
func NewTaobaoTbkOrderRefundGetRequest() *TaobaoTbkOrderRefundGetAPIRequest {
	return &TaobaoTbkOrderRefundGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkOrderRefundGetAPIRequest) Reset() {
	r._publisherRefundOrderQueryOption = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkOrderRefundGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.order.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkOrderRefundGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkOrderRefundGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPublisherRefundOrderQueryOption is PublisherRefundOrderQueryOption Setter
// 全量维权订单查询入参
func (r *TaobaoTbkOrderRefundGetAPIRequest) SetPublisherRefundOrderQueryOption(_publisherRefundOrderQueryOption *PublisherRefundOrderQueryOption) error {
	r._publisherRefundOrderQueryOption = _publisherRefundOrderQueryOption
	r.Set("publisher_refund_order_query_option", _publisherRefundOrderQueryOption)
	return nil
}

// GetPublisherRefundOrderQueryOption PublisherRefundOrderQueryOption Getter
func (r TaobaoTbkOrderRefundGetAPIRequest) GetPublisherRefundOrderQueryOption() *PublisherRefundOrderQueryOption {
	return r._publisherRefundOrderQueryOption
}

var poolTaobaoTbkOrderRefundGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkOrderRefundGetRequest()
	},
}

// GetTaobaoTbkOrderRefundGetRequest 从 sync.Pool 获取 TaobaoTbkOrderRefundGetAPIRequest
func GetTaobaoTbkOrderRefundGetAPIRequest() *TaobaoTbkOrderRefundGetAPIRequest {
	return poolTaobaoTbkOrderRefundGetAPIRequest.Get().(*TaobaoTbkOrderRefundGetAPIRequest)
}

// ReleaseTaobaoTbkOrderRefundGetAPIRequest 将 TaobaoTbkOrderRefundGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkOrderRefundGetAPIRequest(v *TaobaoTbkOrderRefundGetAPIRequest) {
	v.Reset()
	poolTaobaoTbkOrderRefundGetAPIRequest.Put(v)
}
