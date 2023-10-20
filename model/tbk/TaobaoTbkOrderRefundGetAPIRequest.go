package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkorderrefundgetAPIRequest 淘宝客-推广者-全量维权退款订单查询 API请求
// taobao.tbk.order.refund.get
//
// 淘宝客账户下全量维权退款订单查询
type TaobaotbkorderrefundgetAPIRequest struct {
	model.Params
	// 全量维权订单查询入参
	_publisherRefundOrderQueryOption *PublisherRefundOrderQueryOption
}

// NewTaobaotbkorderrefundgetRequest 初始化TaobaotbkorderrefundgetAPIRequest对象
func NewTaobaotbkorderrefundgetRequest() *TaobaotbkorderrefundgetAPIRequest {
	return &TaobaotbkorderrefundgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkorderrefundgetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.order.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkorderrefundgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkorderrefundgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPublisherRefundOrderQueryOption is PublisherRefundOrderQueryOption Setter
// 全量维权订单查询入参
func (r *TaobaotbkorderrefundgetAPIRequest) SetPublisherRefundOrderQueryOption(_publisherRefundOrderQueryOption *PublisherRefundOrderQueryOption) error {
	r._publisherRefundOrderQueryOption = _publisherRefundOrderQueryOption
	r.Set("publisher_refund_order_query_option", _publisherRefundOrderQueryOption)
	return nil
}

// GetPublisherRefundOrderQueryOption PublisherRefundOrderQueryOption Getter
func (r TaobaotbkorderrefundgetAPIRequest) GetPublisherRefundOrderQueryOption() *PublisherRefundOrderQueryOption {
	return r._publisherRefundOrderQueryOption
}
