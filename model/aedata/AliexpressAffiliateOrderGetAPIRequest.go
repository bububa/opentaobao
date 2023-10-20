package aedata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressaffiliateordergetAPIRequest AE流量订单详情获取API API请求
// aliexpress.affiliate.order.get
//
// 联盟推广订单效果获取API
type AliexpressaffiliateordergetAPIRequest struct {
	model.Params
	// 安全签名
	_appSignature string
	// 返回的字段列表
	_fields string
	// 订单ID列表，以逗号分隔，当前只支持子订单ID查询
	_orderIds string
}

// NewAliexpressaffiliateordergetRequest 初始化AliexpressaffiliateordergetAPIRequest对象
func NewAliexpressaffiliateordergetRequest() *AliexpressaffiliateordergetAPIRequest {
	return &AliexpressaffiliateordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressaffiliateordergetAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressaffiliateordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressaffiliateordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppSignature is AppSignature Setter
// 安全签名
func (r *AliexpressaffiliateordergetAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// GetAppSignature AppSignature Getter
func (r AliexpressaffiliateordergetAPIRequest) GetAppSignature() string {
	return r._appSignature
}

// SetFields is Fields Setter
// 返回的字段列表
func (r *AliexpressaffiliateordergetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r AliexpressaffiliateordergetAPIRequest) GetFields() string {
	return r._fields
}

// SetOrderIds is OrderIds Setter
// 订单ID列表，以逗号分隔，当前只支持子订单ID查询
func (r *AliexpressaffiliateordergetAPIRequest) SetOrderIds(_orderIds string) error {
	r._orderIds = _orderIds
	r.Set("order_ids", _orderIds)
	return nil
}

// GetOrderIds OrderIds Getter
func (r AliexpressaffiliateordergetAPIRequest) GetOrderIds() string {
	return r._orderIds
}
