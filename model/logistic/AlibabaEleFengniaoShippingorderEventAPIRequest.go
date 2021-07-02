package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoShippingorderEventAPIRequest 查询运单事件信息 API请求
// alibaba.ele.fengniao.shippingorder.event
//
// 查询运单事件信息
type AlibabaEleFengniaoShippingorderEventAPIRequest struct {
	model.Params
	// appid
	_appId string
	// 外部订单号
	_partnerOrderCode string
}

// NewAlibabaEleFengniaoShippingorderEventRequest 初始化AlibabaEleFengniaoShippingorderEventAPIRequest对象
func NewAlibabaEleFengniaoShippingorderEventRequest() *AlibabaEleFengniaoShippingorderEventAPIRequest {
	return &AlibabaEleFengniaoShippingorderEventAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoShippingorderEventAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.shippingorder.event"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoShippingorderEventAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AppId Setter
// appid
func (r *AlibabaEleFengniaoShippingorderEventAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// Get AppId Getter
func (r AlibabaEleFengniaoShippingorderEventAPIRequest) GetAppId() string {
	return r._appId
}

// Set is PartnerOrderCode Setter
// 外部订单号
func (r *AlibabaEleFengniaoShippingorderEventAPIRequest) SetPartnerOrderCode(_partnerOrderCode string) error {
	r._partnerOrderCode = _partnerOrderCode
	r.Set("partner_order_code", _partnerOrderCode)
	return nil
}

// Get PartnerOrderCode Getter
func (r AlibabaEleFengniaoShippingorderEventAPIRequest) GetPartnerOrderCode() string {
	return r._partnerOrderCode
}
