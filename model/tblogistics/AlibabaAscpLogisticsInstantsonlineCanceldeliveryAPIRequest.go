package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascplogisticsinstantsonlinecanceldeliveryAPIRequest 同城配送在线下单取消下单取消呼叫的运力 API请求
// alibaba.ascp.logistics.instantsonline.canceldelivery
//
// 同城配送在线下单取消下单取消呼叫的运力
type AlibabaascplogisticsinstantsonlinecanceldeliveryAPIRequest struct {
	model.Params
	// 类型
	_bizType string
	// ERP单号
	_outOrderId string
}

// NewAlibabaascplogisticsinstantsonlinecanceldeliveryRequest 初始化AlibabaascplogisticsinstantsonlinecanceldeliveryAPIRequest对象
func NewAlibabaascplogisticsinstantsonlinecanceldeliveryRequest() *AlibabaascplogisticsinstantsonlinecanceldeliveryAPIRequest {
	return &AlibabaascplogisticsinstantsonlinecanceldeliveryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascplogisticsinstantsonlinecanceldeliveryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.instantsonline.canceldelivery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascplogisticsinstantsonlinecanceldeliveryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascplogisticsinstantsonlinecanceldeliveryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizType is BizType Setter
// 类型
func (r *AlibabaascplogisticsinstantsonlinecanceldeliveryAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlibabaascplogisticsinstantsonlinecanceldeliveryAPIRequest) GetBizType() string {
	return r._bizType
}

// SetOutOrderId is OutOrderId Setter
// ERP单号
func (r *AlibabaascplogisticsinstantsonlinecanceldeliveryAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabaascplogisticsinstantsonlinecanceldeliveryAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}
