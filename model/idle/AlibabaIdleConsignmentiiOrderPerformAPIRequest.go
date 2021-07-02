package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleConsignmentiiOrderPerformAPIRequest 寄卖V2订单履约 API请求
// alibaba.idle.consignmentii.order.perform
//
// 寄卖V2订单履约，服务商同步订单信息，驱动交易流转
type AlibabaIdleConsignmentiiOrderPerformAPIRequest struct {
	model.Params
	// 同步参数
	_consignmentV2OrderSynDto *ConsignmentV2OrderSynDto
}

// NewAlibabaIdleConsignmentiiOrderPerformRequest 初始化AlibabaIdleConsignmentiiOrderPerformAPIRequest对象
func NewAlibabaIdleConsignmentiiOrderPerformRequest() *AlibabaIdleConsignmentiiOrderPerformAPIRequest {
	return &AlibabaIdleConsignmentiiOrderPerformAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleConsignmentiiOrderPerformAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.consignmentii.order.perform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleConsignmentiiOrderPerformAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ConsignmentV2OrderSynDto Setter
// 同步参数
func (r *AlibabaIdleConsignmentiiOrderPerformAPIRequest) SetConsignmentV2OrderSynDto(_consignmentV2OrderSynDto *ConsignmentV2OrderSynDto) error {
	r._consignmentV2OrderSynDto = _consignmentV2OrderSynDto
	r.Set("consignment_v2_order_syn_dto", _consignmentV2OrderSynDto)
	return nil
}

// Get ConsignmentV2OrderSynDto Getter
func (r AlibabaIdleConsignmentiiOrderPerformAPIRequest) GetConsignmentV2OrderSynDto() *ConsignmentV2OrderSynDto {
	return r._consignmentV2OrderSynDto
}
