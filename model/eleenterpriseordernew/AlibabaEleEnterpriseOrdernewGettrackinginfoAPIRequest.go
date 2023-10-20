package eleenterpriseordernew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeleenterpriseordernewgettrackinginfoAPIRequest 订单配送信息跟踪 API请求
// alibaba.ele.enterprise.ordernew.gettrackinginfo
//
// 订单配送信息跟踪
type AlibabaeleenterpriseordernewgettrackinginfoAPIRequest struct {
	model.Params
	// 饿了么订单ID
	_orderId string
	// 用户手机号
	_phone string
}

// NewAlibabaeleenterpriseordernewgettrackinginfoRequest 初始化AlibabaeleenterpriseordernewgettrackinginfoAPIRequest对象
func NewAlibabaeleenterpriseordernewgettrackinginfoRequest() *AlibabaeleenterpriseordernewgettrackinginfoAPIRequest {
	return &AlibabaeleenterpriseordernewgettrackinginfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeleenterpriseordernewgettrackinginfoAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.ordernew.gettrackinginfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeleenterpriseordernewgettrackinginfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeleenterpriseordernewgettrackinginfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 饿了么订单ID
func (r *AlibabaeleenterpriseordernewgettrackinginfoAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaeleenterpriseordernewgettrackinginfoAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetPhone is Phone Setter
// 用户手机号
func (r *AlibabaeleenterpriseordernewgettrackinginfoAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaeleenterpriseordernewgettrackinginfoAPIRequest) GetPhone() string {
	return r._phone
}
