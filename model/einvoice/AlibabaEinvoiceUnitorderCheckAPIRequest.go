package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceunitordercheckAPIRequest 服务商订购单上传核对 API请求
// alibaba.einvoice.unitorder.check
//
// 开票服务商回传收到的订购单用于电子发票平台核对
type AlibabaeinvoiceunitordercheckAPIRequest struct {
	model.Params
	// 订购单列表
	_orders []SimpleUnitOrder
	// 开始时间,来自于查询消息
	_begin string
	// 结束时间,来自于查询消息
	_end string
}

// NewAlibabaeinvoiceunitordercheckRequest 初始化AlibabaeinvoiceunitordercheckAPIRequest对象
func NewAlibabaeinvoiceunitordercheckRequest() *AlibabaeinvoiceunitordercheckAPIRequest {
	return &AlibabaeinvoiceunitordercheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceunitordercheckAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.unitorder.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceunitordercheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceunitordercheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrders is Orders Setter
// 订购单列表
func (r *AlibabaeinvoiceunitordercheckAPIRequest) SetOrders(_orders []SimpleUnitOrder) error {
	r._orders = _orders
	r.Set("orders", _orders)
	return nil
}

// GetOrders Orders Getter
func (r AlibabaeinvoiceunitordercheckAPIRequest) GetOrders() []SimpleUnitOrder {
	return r._orders
}

// SetBegin is Begin Setter
// 开始时间,来自于查询消息
func (r *AlibabaeinvoiceunitordercheckAPIRequest) SetBegin(_begin string) error {
	r._begin = _begin
	r.Set("begin", _begin)
	return nil
}

// GetBegin Begin Getter
func (r AlibabaeinvoiceunitordercheckAPIRequest) GetBegin() string {
	return r._begin
}

// SetEnd is End Setter
// 结束时间,来自于查询消息
func (r *AlibabaeinvoiceunitordercheckAPIRequest) SetEnd(_end string) error {
	r._end = _end
	r.Set("end", _end)
	return nil
}

// GetEnd End Getter
func (r AlibabaeinvoiceunitordercheckAPIRequest) GetEnd() string {
	return r._end
}
