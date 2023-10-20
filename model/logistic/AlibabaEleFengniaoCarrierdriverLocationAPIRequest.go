package logistic

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleFengniaoCarrierdriverLocationAPIRequest) Reset() {
	r._appId = ""
	r._partnerOrderCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoCarrierdriverLocationAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.carrierdriver.location"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoCarrierdriverLocationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleFengniaoCarrierdriverLocationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppId is AppId Setter
// appid
func (r *AlibabaEleFengniaoCarrierdriverLocationAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabaEleFengniaoCarrierdriverLocationAPIRequest) GetAppId() string {
	return r._appId
}

// SetPartnerOrderCode is PartnerOrderCode Setter
// 外部订单号
func (r *AlibabaEleFengniaoCarrierdriverLocationAPIRequest) SetPartnerOrderCode(_partnerOrderCode string) error {
	r._partnerOrderCode = _partnerOrderCode
	r.Set("partner_order_code", _partnerOrderCode)
	return nil
}

// GetPartnerOrderCode PartnerOrderCode Getter
func (r AlibabaEleFengniaoCarrierdriverLocationAPIRequest) GetPartnerOrderCode() string {
	return r._partnerOrderCode
}

var poolAlibabaEleFengniaoCarrierdriverLocationAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleFengniaoCarrierdriverLocationRequest()
	},
}

// GetAlibabaEleFengniaoCarrierdriverLocationRequest 从 sync.Pool 获取 AlibabaEleFengniaoCarrierdriverLocationAPIRequest
func GetAlibabaEleFengniaoCarrierdriverLocationAPIRequest() *AlibabaEleFengniaoCarrierdriverLocationAPIRequest {
	return poolAlibabaEleFengniaoCarrierdriverLocationAPIRequest.Get().(*AlibabaEleFengniaoCarrierdriverLocationAPIRequest)
}

// ReleaseAlibabaEleFengniaoCarrierdriverLocationAPIRequest 将 AlibabaEleFengniaoCarrierdriverLocationAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleFengniaoCarrierdriverLocationAPIRequest(v *AlibabaEleFengniaoCarrierdriverLocationAPIRequest) {
	v.Reset()
	poolAlibabaEleFengniaoCarrierdriverLocationAPIRequest.Put(v)
}
