package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaelefengniaoshippingordereventAPIRequest 查询运单事件信息 API请求
// alibaba.ele.fengniao.shippingorder.event
//
// 查询运单事件信息
type AlibabaelefengniaoshippingordereventAPIRequest struct {
	model.Params
	// appid
	_appId string
	// 外部订单号
	_partnerOrderCode string
}

// NewAlibabaelefengniaoshippingordereventRequest 初始化AlibabaelefengniaoshippingordereventAPIRequest对象
func NewAlibabaelefengniaoshippingordereventRequest() *AlibabaelefengniaoshippingordereventAPIRequest {
	return &AlibabaelefengniaoshippingordereventAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaelefengniaoshippingordereventAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.shippingorder.event"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaelefengniaoshippingordereventAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaelefengniaoshippingordereventAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppId is AppId Setter
// appid
func (r *AlibabaelefengniaoshippingordereventAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabaelefengniaoshippingordereventAPIRequest) GetAppId() string {
	return r._appId
}

// SetPartnerOrderCode is PartnerOrderCode Setter
// 外部订单号
func (r *AlibabaelefengniaoshippingordereventAPIRequest) SetPartnerOrderCode(_partnerOrderCode string) error {
	r._partnerOrderCode = _partnerOrderCode
	r.Set("partner_order_code", _partnerOrderCode)
	return nil
}

// GetPartnerOrderCode PartnerOrderCode Getter
func (r AlibabaelefengniaoshippingordereventAPIRequest) GetPartnerOrderCode() string {
	return r._partnerOrderCode
}
