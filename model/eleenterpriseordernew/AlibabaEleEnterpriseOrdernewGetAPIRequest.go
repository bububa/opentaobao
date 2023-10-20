package eleenterpriseordernew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeleenterpriseordernewgetAPIRequest 查询订单详情 API请求
// alibaba.ele.enterprise.ordernew.get
//
// 查询订单详情
type AlibabaeleenterpriseordernewgetAPIRequest struct {
	model.Params
	// 饿了么订单ID
	_orderId string
	// 电话号码
	_phone string
}

// NewAlibabaeleenterpriseordernewgetRequest 初始化AlibabaeleenterpriseordernewgetAPIRequest对象
func NewAlibabaeleenterpriseordernewgetRequest() *AlibabaeleenterpriseordernewgetAPIRequest {
	return &AlibabaeleenterpriseordernewgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeleenterpriseordernewgetAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.ordernew.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeleenterpriseordernewgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeleenterpriseordernewgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 饿了么订单ID
func (r *AlibabaeleenterpriseordernewgetAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaeleenterpriseordernewgetAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetPhone is Phone Setter
// 电话号码
func (r *AlibabaeleenterpriseordernewgetAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaeleenterpriseordernewgetAPIRequest) GetPhone() string {
	return r._phone
}
