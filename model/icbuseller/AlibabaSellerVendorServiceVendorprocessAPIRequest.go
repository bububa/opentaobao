package icbuseller

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasellervendorservicevendorprocessAPIRequest 服务商客户关联信息 API请求
// alibaba.seller.vendor.service.vendorprocess
//
// 服务商客户关联信息
type AlibabasellervendorservicevendorprocessAPIRequest struct {
	model.Params
	// order_num
	_orderNum string
}

// NewAlibabasellervendorservicevendorprocessRequest 初始化AlibabasellervendorservicevendorprocessAPIRequest对象
func NewAlibabasellervendorservicevendorprocessRequest() *AlibabasellervendorservicevendorprocessAPIRequest {
	return &AlibabasellervendorservicevendorprocessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasellervendorservicevendorprocessAPIRequest) GetApiMethodName() string {
	return "alibaba.seller.vendor.service.vendorprocess"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasellervendorservicevendorprocessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasellervendorservicevendorprocessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderNum is OrderNum Setter
// order_num
func (r *AlibabasellervendorservicevendorprocessAPIRequest) SetOrderNum(_orderNum string) error {
	r._orderNum = _orderNum
	r.Set("order_num", _orderNum)
	return nil
}

// GetOrderNum OrderNum Getter
func (r AlibabasellervendorservicevendorprocessAPIRequest) GetOrderNum() string {
	return r._orderNum
}
