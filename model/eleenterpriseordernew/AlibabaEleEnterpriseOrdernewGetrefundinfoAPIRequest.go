package eleenterpriseordernew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseOrdernewGetrefundinfoAPIRequest 退单和申诉 API请求
// alibaba.ele.enterprise.ordernew.getrefundinfo
//
// 退单和申诉
type AlibabaEleEnterpriseOrdernewGetrefundinfoAPIRequest struct {
	model.Params
	// 饿了么订单ID
	_orderId string
}

// NewAlibabaEleEnterpriseOrdernewGetrefundinfoRequest 初始化AlibabaEleEnterpriseOrdernewGetrefundinfoAPIRequest对象
func NewAlibabaEleEnterpriseOrdernewGetrefundinfoRequest() *AlibabaEleEnterpriseOrdernewGetrefundinfoAPIRequest {
	return &AlibabaEleEnterpriseOrdernewGetrefundinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseOrdernewGetrefundinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.ordernew.getrefundinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseOrdernewGetrefundinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleEnterpriseOrdernewGetrefundinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 饿了么订单ID
func (r *AlibabaEleEnterpriseOrdernewGetrefundinfoAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaEleEnterpriseOrdernewGetrefundinfoAPIRequest) GetOrderId() string {
	return r._orderId
}
