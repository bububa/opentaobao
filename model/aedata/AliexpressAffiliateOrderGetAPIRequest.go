package aedata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateOrderGetAPIRequest AE流量订单详情获取API API请求
// aliexpress.affiliate.order.get
//
// 联盟推广订单效果获取API
type AliexpressAffiliateOrderGetAPIRequest struct {
	model.Params
	// 安全签名
	_appSignature string
	// 返回的字段列表
	_fields string
	// 订单ID列表，以逗号分隔，当前只支持子订单ID查询
	_orderIds string
}

// NewAliexpressAffiliateOrderGetRequest 初始化AliexpressAffiliateOrderGetAPIRequest对象
func NewAliexpressAffiliateOrderGetRequest() *AliexpressAffiliateOrderGetAPIRequest {
	return &AliexpressAffiliateOrderGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressAffiliateOrderGetAPIRequest) Reset() {
	r._appSignature = ""
	r._fields = ""
	r._orderIds = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateOrderGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAffiliateOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppSignature is AppSignature Setter
// 安全签名
func (r *AliexpressAffiliateOrderGetAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// GetAppSignature AppSignature Getter
func (r AliexpressAffiliateOrderGetAPIRequest) GetAppSignature() string {
	return r._appSignature
}

// SetFields is Fields Setter
// 返回的字段列表
func (r *AliexpressAffiliateOrderGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r AliexpressAffiliateOrderGetAPIRequest) GetFields() string {
	return r._fields
}

// SetOrderIds is OrderIds Setter
// 订单ID列表，以逗号分隔，当前只支持子订单ID查询
func (r *AliexpressAffiliateOrderGetAPIRequest) SetOrderIds(_orderIds string) error {
	r._orderIds = _orderIds
	r.Set("order_ids", _orderIds)
	return nil
}

// GetOrderIds OrderIds Getter
func (r AliexpressAffiliateOrderGetAPIRequest) GetOrderIds() string {
	return r._orderIds
}

var poolAliexpressAffiliateOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressAffiliateOrderGetRequest()
	},
}

// GetAliexpressAffiliateOrderGetRequest 从 sync.Pool 获取 AliexpressAffiliateOrderGetAPIRequest
func GetAliexpressAffiliateOrderGetAPIRequest() *AliexpressAffiliateOrderGetAPIRequest {
	return poolAliexpressAffiliateOrderGetAPIRequest.Get().(*AliexpressAffiliateOrderGetAPIRequest)
}

// ReleaseAliexpressAffiliateOrderGetAPIRequest 将 AliexpressAffiliateOrderGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressAffiliateOrderGetAPIRequest(v *AliexpressAffiliateOrderGetAPIRequest) {
	v.Reset()
	poolAliexpressAffiliateOrderGetAPIRequest.Put(v)
}
