package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchannelsalesorderconfirmAPIRequest 渠道销售单确认收货 API请求
// alibaba.ascp.channel.sales.order.confirm
//
// 渠道销售单确认收货
type AlibabaascpchannelsalesorderconfirmAPIRequest struct {
	model.Params
	// 请求参数
	_confirmOrderRequest *ExternalConfirmSalesOrderRequest
}

// NewAlibabaascpchannelsalesorderconfirmRequest 初始化AlibabaascpchannelsalesorderconfirmAPIRequest对象
func NewAlibabaascpchannelsalesorderconfirmRequest() *AlibabaascpchannelsalesorderconfirmAPIRequest {
	return &AlibabaascpchannelsalesorderconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchannelsalesorderconfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.sales.order.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchannelsalesorderconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchannelsalesorderconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConfirmOrderRequest is ConfirmOrderRequest Setter
// 请求参数
func (r *AlibabaascpchannelsalesorderconfirmAPIRequest) SetConfirmOrderRequest(_confirmOrderRequest *ExternalConfirmSalesOrderRequest) error {
	r._confirmOrderRequest = _confirmOrderRequest
	r.Set("confirm_order_request", _confirmOrderRequest)
	return nil
}

// GetConfirmOrderRequest ConfirmOrderRequest Getter
func (r AlibabaascpchannelsalesorderconfirmAPIRequest) GetConfirmOrderRequest() *ExternalConfirmSalesOrderRequest {
	return r._confirmOrderRequest
}
