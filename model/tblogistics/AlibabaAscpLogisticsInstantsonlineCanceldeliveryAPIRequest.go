package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIRequest 同城配送在线下单取消下单取消呼叫的运力 API请求
// alibaba.ascp.logistics.instantsonline.canceldelivery
//
// 同城配送在线下单取消下单取消呼叫的运力
type AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIRequest struct {
	model.Params
	// 类型
	_bizType string
	// ERP单号
	_outOrderId string
}

// NewAlibabaAscpLogisticsInstantsonlineCanceldeliveryRequest 初始化AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIRequest对象
func NewAlibabaAscpLogisticsInstantsonlineCanceldeliveryRequest() *AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIRequest {
	return &AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.instantsonline.canceldelivery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizType is BizType Setter
// 类型
func (r *AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIRequest) GetBizType() string {
	return r._bizType
}

// SetOutOrderId is OutOrderId Setter
// ERP单号
func (r *AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}
