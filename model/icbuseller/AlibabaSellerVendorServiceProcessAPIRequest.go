package icbuseller

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasellervendorserviceprocessAPIRequest 服务商客户关联信息 API请求
// alibaba.seller.vendor.service.process
//
// 服务商客户关联信息
type AlibabasellervendorserviceprocessAPIRequest struct {
	model.Params
	// order_num
	_orderNum string
}

// NewAlibabasellervendorserviceprocessRequest 初始化AlibabasellervendorserviceprocessAPIRequest对象
func NewAlibabasellervendorserviceprocessRequest() *AlibabasellervendorserviceprocessAPIRequest {
	return &AlibabasellervendorserviceprocessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasellervendorserviceprocessAPIRequest) GetApiMethodName() string {
	return "alibaba.seller.vendor.service.process"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasellervendorserviceprocessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasellervendorserviceprocessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderNum is OrderNum Setter
// order_num
func (r *AlibabasellervendorserviceprocessAPIRequest) SetOrderNum(_orderNum string) error {
	r._orderNum = _orderNum
	r.Set("order_num", _orderNum)
	return nil
}

// GetOrderNum OrderNum Getter
func (r AlibabasellervendorserviceprocessAPIRequest) GetOrderNum() string {
	return r._orderNum
}
