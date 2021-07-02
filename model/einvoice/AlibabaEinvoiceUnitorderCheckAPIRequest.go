package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceUnitorderCheckAPIRequest 服务商订购单上传核对 API请求
// alibaba.einvoice.unitorder.check
//
// 开票服务商回传收到的订购单用于电子发票平台核对
type AlibabaEinvoiceUnitorderCheckAPIRequest struct {
	model.Params
	// 订购单列表
	_orders []SimpleUnitOrder
	// 开始时间,来自于查询消息
	_begin string
	// 结束时间,来自于查询消息
	_end string
}

// NewAlibabaEinvoiceUnitorderCheckRequest 初始化AlibabaEinvoiceUnitorderCheckAPIRequest对象
func NewAlibabaEinvoiceUnitorderCheckRequest() *AlibabaEinvoiceUnitorderCheckAPIRequest {
	return &AlibabaEinvoiceUnitorderCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceUnitorderCheckAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.unitorder.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceUnitorderCheckAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Orders Setter
// 订购单列表
func (r *AlibabaEinvoiceUnitorderCheckAPIRequest) SetOrders(_orders []SimpleUnitOrder) error {
	r._orders = _orders
	r.Set("orders", _orders)
	return nil
}

// Get Orders Getter
func (r AlibabaEinvoiceUnitorderCheckAPIRequest) GetOrders() []SimpleUnitOrder {
	return r._orders
}

// Set is Begin Setter
// 开始时间,来自于查询消息
func (r *AlibabaEinvoiceUnitorderCheckAPIRequest) SetBegin(_begin string) error {
	r._begin = _begin
	r.Set("begin", _begin)
	return nil
}

// Get Begin Getter
func (r AlibabaEinvoiceUnitorderCheckAPIRequest) GetBegin() string {
	return r._begin
}

// Set is End Setter
// 结束时间,来自于查询消息
func (r *AlibabaEinvoiceUnitorderCheckAPIRequest) SetEnd(_end string) error {
	r._end = _end
	r.Set("end", _end)
	return nil
}

// Get End Getter
func (r AlibabaEinvoiceUnitorderCheckAPIRequest) GetEnd() string {
	return r._end
}
