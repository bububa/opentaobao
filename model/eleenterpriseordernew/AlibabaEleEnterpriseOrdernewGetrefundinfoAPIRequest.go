package eleenterpriseordernew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeleenterpriseordernewgetrefundinfoAPIRequest 退单和申诉 API请求
// alibaba.ele.enterprise.ordernew.getrefundinfo
//
// 退单和申诉
type AlibabaeleenterpriseordernewgetrefundinfoAPIRequest struct {
	model.Params
	// 饿了么订单ID
	_orderId string
}

// NewAlibabaeleenterpriseordernewgetrefundinfoRequest 初始化AlibabaeleenterpriseordernewgetrefundinfoAPIRequest对象
func NewAlibabaeleenterpriseordernewgetrefundinfoRequest() *AlibabaeleenterpriseordernewgetrefundinfoAPIRequest {
	return &AlibabaeleenterpriseordernewgetrefundinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeleenterpriseordernewgetrefundinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.ordernew.getrefundinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeleenterpriseordernewgetrefundinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeleenterpriseordernewgetrefundinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 饿了么订单ID
func (r *AlibabaeleenterpriseordernewgetrefundinfoAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaeleenterpriseordernewgetrefundinfoAPIRequest) GetOrderId() string {
	return r._orderId
}
