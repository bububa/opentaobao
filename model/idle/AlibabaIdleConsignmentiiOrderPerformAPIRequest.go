package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleconsignmentiiorderperformAPIRequest 寄卖V2订单履约 API请求
// alibaba.idle.consignmentii.order.perform
//
// 寄卖V2订单履约，服务商同步订单信息，驱动交易流转
type AlibabaidleconsignmentiiorderperformAPIRequest struct {
	model.Params
	// 同步参数
	_consignmentV2OrderSynDto *ConsignmentV2orderSynDto
}

// NewAlibabaidleconsignmentiiorderperformRequest 初始化AlibabaidleconsignmentiiorderperformAPIRequest对象
func NewAlibabaidleconsignmentiiorderperformRequest() *AlibabaidleconsignmentiiorderperformAPIRequest {
	return &AlibabaidleconsignmentiiorderperformAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleconsignmentiiorderperformAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.consignmentii.order.perform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleconsignmentiiorderperformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleconsignmentiiorderperformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConsignmentV2OrderSynDto is ConsignmentV2OrderSynDto Setter
// 同步参数
func (r *AlibabaidleconsignmentiiorderperformAPIRequest) SetConsignmentV2OrderSynDto(_consignmentV2OrderSynDto *ConsignmentV2orderSynDto) error {
	r._consignmentV2OrderSynDto = _consignmentV2OrderSynDto
	r.Set("consignment_v2_order_syn_dto", _consignmentV2OrderSynDto)
	return nil
}

// GetConsignmentV2OrderSynDto ConsignmentV2OrderSynDto Getter
func (r AlibabaidleconsignmentiiorderperformAPIRequest) GetConsignmentV2OrderSynDto() *ConsignmentV2orderSynDto {
	return r._consignmentV2OrderSynDto
}
