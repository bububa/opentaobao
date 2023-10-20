package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaelefengniaocarrierdriverlocationAPIRequest 查询骑手当前位置 API请求
// alibaba.ele.fengniao.carrierdriver.location
//
// 查询骑手当前位置
type AlibabaelefengniaocarrierdriverlocationAPIRequest struct {
	model.Params
	// appid
	_appId string
	// 外部订单号
	_partnerOrderCode string
}

// NewAlibabaelefengniaocarrierdriverlocationRequest 初始化AlibabaelefengniaocarrierdriverlocationAPIRequest对象
func NewAlibabaelefengniaocarrierdriverlocationRequest() *AlibabaelefengniaocarrierdriverlocationAPIRequest {
	return &AlibabaelefengniaocarrierdriverlocationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaelefengniaocarrierdriverlocationAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.carrierdriver.location"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaelefengniaocarrierdriverlocationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaelefengniaocarrierdriverlocationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppId is AppId Setter
// appid
func (r *AlibabaelefengniaocarrierdriverlocationAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabaelefengniaocarrierdriverlocationAPIRequest) GetAppId() string {
	return r._appId
}

// SetPartnerOrderCode is PartnerOrderCode Setter
// 外部订单号
func (r *AlibabaelefengniaocarrierdriverlocationAPIRequest) SetPartnerOrderCode(_partnerOrderCode string) error {
	r._partnerOrderCode = _partnerOrderCode
	r.Set("partner_order_code", _partnerOrderCode)
	return nil
}

// GetPartnerOrderCode PartnerOrderCode Getter
func (r AlibabaelefengniaocarrierdriverlocationAPIRequest) GetPartnerOrderCode() string {
	return r._partnerOrderCode
}
