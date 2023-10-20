package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchannelmainrefundcreateAPIRequest 淘外分销逆向创单（未发货整单退） API请求
// alibaba.ascp.channel.main.refund.create
//
// 淘外分销解决方案--订单--逆向创单（未发货整单退）
type AlibabaascpchannelmainrefundcreateAPIRequest struct {
	model.Params
	// 逆向单创建请求
	_refundCreateRequest *ExternalCreateRefundOrderRequest
}

// NewAlibabaascpchannelmainrefundcreateRequest 初始化AlibabaascpchannelmainrefundcreateAPIRequest对象
func NewAlibabaascpchannelmainrefundcreateRequest() *AlibabaascpchannelmainrefundcreateAPIRequest {
	return &AlibabaascpchannelmainrefundcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchannelmainrefundcreateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.main.refund.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchannelmainrefundcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchannelmainrefundcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundCreateRequest is RefundCreateRequest Setter
// 逆向单创建请求
func (r *AlibabaascpchannelmainrefundcreateAPIRequest) SetRefundCreateRequest(_refundCreateRequest *ExternalCreateRefundOrderRequest) error {
	r._refundCreateRequest = _refundCreateRequest
	r.Set("refund_create_request", _refundCreateRequest)
	return nil
}

// GetRefundCreateRequest RefundCreateRequest Getter
func (r AlibabaascpchannelmainrefundcreateAPIRequest) GetRefundCreateRequest() *ExternalCreateRefundOrderRequest {
	return r._refundCreateRequest
}
