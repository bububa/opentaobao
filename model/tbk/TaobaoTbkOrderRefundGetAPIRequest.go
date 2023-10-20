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
	_publisherrefundorderqueryoption *PublisherRefundOrderQueryOption
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

// SetPublisherrefundorderqueryoption is Publisherrefundorderqueryoption Setter
// 全量维权订单查询入参
func (r *TaobaotbkorderrefundgetAPIRequest) SetPublisherrefundorderqueryoption(_publisherrefundorderqueryoption *PublisherRefundOrderQueryOption) error {
	r._publisherrefundorderqueryoption = _publisherrefundorderqueryoption
	r.Set("publisher_refund_order_query_option", _publisherrefundorderqueryoption)
	return nil
}

// GetPublisherrefundorderqueryoption Publisherrefundorderqueryoption Getter
func (r TaobaotbkorderrefundgetAPIRequest) GetPublisherrefundorderqueryoption() *PublisherRefundOrderQueryOption {
	return r._publisherrefundorderqueryoption
}
