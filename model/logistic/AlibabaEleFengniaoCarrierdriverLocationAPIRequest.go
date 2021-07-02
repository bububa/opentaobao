package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoCarrierdriverLocationAPIRequest 查询骑手当前位置 API请求
// alibaba.ele.fengniao.carrierdriver.location
//
// 查询骑手当前位置
type AlibabaEleFengniaoCarrierdriverLocationAPIRequest struct {
	model.Params
	// appid
	_appId string
	// 外部订单号
	_partnerOrderCode string
}

// NewAlibabaEleFengniaoCarrierdriverLocationRequest 初始化AlibabaEleFengniaoCarrierdriverLocationAPIRequest对象
func NewAlibabaEleFengniaoCarrierdriverLocationRequest() *AlibabaEleFengniaoCarrierdriverLocationAPIRequest {
	return &AlibabaEleFengniaoCarrierdriverLocationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoCarrierdriverLocationAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.carrierdriver.location"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoCarrierdriverLocationAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AppId Setter
// appid
func (r *AlibabaEleFengniaoCarrierdriverLocationAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// Get AppId Getter
func (r AlibabaEleFengniaoCarrierdriverLocationAPIRequest) GetAppId() string {
	return r._appId
}

// Set is PartnerOrderCode Setter
// 外部订单号
func (r *AlibabaEleFengniaoCarrierdriverLocationAPIRequest) SetPartnerOrderCode(_partnerOrderCode string) error {
	r._partnerOrderCode = _partnerOrderCode
	r.Set("partner_order_code", _partnerOrderCode)
	return nil
}

// Get PartnerOrderCode Getter
func (r AlibabaEleFengniaoCarrierdriverLocationAPIRequest) GetPartnerOrderCode() string {
	return r._partnerOrderCode
}
