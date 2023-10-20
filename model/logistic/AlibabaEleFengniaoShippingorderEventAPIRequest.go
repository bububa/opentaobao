package logistic

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleFengniaoShippingorderEventAPIRequest) Reset() {
	r._appId = ""
	r._partnerOrderCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoShippingorderEventAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.shippingorder.event"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoShippingorderEventAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleFengniaoShippingorderEventAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppId is AppId Setter
// appid
func (r *AlibabaEleFengniaoShippingorderEventAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabaEleFengniaoShippingorderEventAPIRequest) GetAppId() string {
	return r._appId
}

// SetPartnerOrderCode is PartnerOrderCode Setter
// 外部订单号
func (r *AlibabaEleFengniaoShippingorderEventAPIRequest) SetPartnerOrderCode(_partnerOrderCode string) error {
	r._partnerOrderCode = _partnerOrderCode
	r.Set("partner_order_code", _partnerOrderCode)
	return nil
}

// GetPartnerOrderCode PartnerOrderCode Getter
func (r AlibabaEleFengniaoShippingorderEventAPIRequest) GetPartnerOrderCode() string {
	return r._partnerOrderCode
}

var poolAlibabaEleFengniaoShippingorderEventAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleFengniaoShippingorderEventRequest()
	},
}

// GetAlibabaEleFengniaoShippingorderEventRequest 从 sync.Pool 获取 AlibabaEleFengniaoShippingorderEventAPIRequest
func GetAlibabaEleFengniaoShippingorderEventAPIRequest() *AlibabaEleFengniaoShippingorderEventAPIRequest {
	return poolAlibabaEleFengniaoShippingorderEventAPIRequest.Get().(*AlibabaEleFengniaoShippingorderEventAPIRequest)
}

// ReleaseAlibabaEleFengniaoShippingorderEventAPIRequest 将 AlibabaEleFengniaoShippingorderEventAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleFengniaoShippingorderEventAPIRequest(v *AlibabaEleFengniaoShippingorderEventAPIRequest) {
	v.Reset()
	poolAlibabaEleFengniaoShippingorderEventAPIRequest.Put(v)
}
